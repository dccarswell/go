// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func collect[T any](seq Seq[T]) []T {
	var got []T
	for v := range seq {
		got = append(got, v)
	}
	return got
}

func collect2[K comparable, V any](seq Seq2[K, V]) ([]K, []V) {
	var keys []K
	var vals []V
	for k, v := range seq {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	return keys, vals
}

func assertIntSlice(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestFilterKeepsMatching(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4, 5, 6})
	even := func(v int) bool { return v%2 == 0 }

	got := collect(Filter(FilterFunc[int](even))(seq))
	assertIntSlice(t, got, []int{2, 4, 6})
}

func TestFilterMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4, 5, 6})
	even := func(v int) bool { return v%2 == 0 }

	got := collect(seq.Filter(even))
	assertIntSlice(t, got, []int{2, 4, 6})
}

func TestFilterEmpty(t *testing.T) {
	got := collect(sliceSeq[int](nil).Filter(func(int) bool { return true }))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestFilterRejectsAll(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})
	got := collect(seq.Filter(func(int) bool { return false }))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestFilterEarlyTermination(t *testing.T) {
	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 100; i++ {
			calls++
			if !yield(i) {
				return
			}
		}
	})

	var got []int
	for v := range source.Filter(func(int) bool { return true }) {
		got = append(got, v)
		if len(got) == 3 {
			break
		}
	}

	assertIntSlice(t, got, []int{1, 2, 3})
	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (early termination not propagated)", calls)
	}
}

func TestFilter2KeepsMatching(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4})
	evenVal := func(k string, v int) bool { return v%2 == 0 }

	keys, vals := collect2(Filter2(FilterFunc2[string, int](evenVal))(seq))
	if len(keys) != 2 || keys[0] != "b" || keys[1] != "d" || vals[0] != 2 || vals[1] != 4 {
		t.Errorf("keys=%v vals=%v, want keys=[b d] vals=[2 4]", keys, vals)
	}
}

func TestFilter2Method(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4})
	evenVal := func(k string, v int) bool { return v%2 == 0 }

	keys, vals := collect2(seq.Filter(evenVal))
	if len(keys) != 2 || keys[0] != "b" || keys[1] != "d" || vals[0] != 2 || vals[1] != 4 {
		t.Errorf("keys=%v vals=%v, want keys=[b d] vals=[2 4]", keys, vals)
	}
}

func TestFilter2Empty(t *testing.T) {
	keys, vals := collect2(sliceSeq2[string, int](nil, nil).Filter(func(string, int) bool { return true }))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty", keys, vals)
	}
}
