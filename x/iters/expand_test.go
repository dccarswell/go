// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestExpand(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})
	repeat := func(v int) Seq[int] { return sliceSeq([]int{v, v}) }

	got := collect(Expand(ExpandFunc[int](repeat))(seq))
	assertIntSlice(t, got, []int{1, 1, 2, 2, 3, 3})
}

func TestExpandMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})
	repeat := func(v int) Seq[int] { return sliceSeq([]int{v, v}) }

	got := collect(seq.Expand(repeat))
	assertIntSlice(t, got, []int{1, 1, 2, 2, 3, 3})
}

func TestExpandEmptySource(t *testing.T) {
	repeat := func(v int) Seq[int] { return sliceSeq([]int{v, v}) }

	got := collect(sliceSeq[int](nil).Expand(repeat))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestExpandSubSequenceCanBeEmpty(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})
	onlyEven := func(v int) Seq[int] {
		if v%2 == 0 {
			return sliceSeq([]int{v})
		}
		return sliceSeq[int](nil)
	}

	got := collect(seq.Expand(onlyEven))
	assertIntSlice(t, got, []int{2})
}

func TestExpandPreservesOrderAcrossSubSequences(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})
	rangeUpTo := func(v int) Seq[int] {
		items := make([]int, v)
		for i := range items {
			items[i] = v*10 + i
		}
		return sliceSeq(items)
	}

	got := collect(seq.Expand(rangeUpTo))
	// v=1 -> [10]; v=2 -> [20 21]; v=3 -> [30 31 32]
	assertIntSlice(t, got, []int{10, 20, 21, 30, 31, 32})
}

func TestExpandEarlyTerminationWithinSubSequence(t *testing.T) {
	outerCalls := 0
	seq := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 100; i++ {
			outerCalls++
			if !yield(i) {
				return
			}
		}
	})
	repeat := func(v int) Seq[int] { return sliceSeq([]int{v, v, v}) }

	var got []int
	for v := range seq.Expand(repeat) {
		got = append(got, v)
		if len(got) == 2 {
			break
		}
	}

	assertIntSlice(t, got, []int{1, 1})
	if outerCalls != 1 {
		t.Errorf("outer source yielded %d times, want 1 (stopping mid-sub-sequence should stop the outer source too)", outerCalls)
	}
}

func TestExpand2(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})
	dup := func(k string, v int) Seq2[string, int] {
		return sliceSeq2([]string{k, k}, []int{v, v})
	}

	keys, vals := collect2(Expand2(ExpandFunc2[string, int](dup))(seq))
	wantKeys := []string{"a", "a", "b", "b"}
	wantVals := []int{1, 1, 2, 2}
	if len(keys) != len(wantKeys) {
		t.Fatalf("keys=%v vals=%v, want keys=%v vals=%v", keys, vals, wantKeys, wantVals)
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] || vals[i] != wantVals[i] {
			t.Errorf("pair[%d] = (%q, %d), want (%q, %d)", i, keys[i], vals[i], wantKeys[i], wantVals[i])
		}
	}
}

func TestExpand2Method(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})
	dup := func(k string, v int) Seq2[string, int] {
		return sliceSeq2([]string{k, k}, []int{v, v})
	}

	keys, vals := collect2(seq.Expand(dup))
	wantKeys := []string{"a", "a", "b", "b"}
	wantVals := []int{1, 1, 2, 2}
	if len(keys) != len(wantKeys) {
		t.Fatalf("keys=%v vals=%v, want keys=%v vals=%v", keys, vals, wantKeys, wantVals)
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] || vals[i] != wantVals[i] {
			t.Errorf("pair[%d] = (%q, %d), want (%q, %d)", i, keys[i], vals[i], wantKeys[i], wantVals[i])
		}
	}
}

func TestExpand2EmptySource(t *testing.T) {
	dup := func(k string, v int) Seq2[string, int] {
		return sliceSeq2([]string{k, k}, []int{v, v})
	}

	keys, vals := collect2(sliceSeq2[string, int](nil, nil).Expand(dup))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty", keys, vals)
	}
}

func TestExpand2SubSequenceCanBeEmpty(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})
	onlyEven := func(k string, v int) Seq2[string, int] {
		if v%2 == 0 {
			return sliceSeq2([]string{k}, []int{v})
		}
		return sliceSeq2[string, int](nil, nil)
	}

	keys, vals := collect2(seq.Expand(onlyEven))
	if len(keys) != 1 || keys[0] != "b" || vals[0] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[b] vals=[2]", keys, vals)
	}
}
