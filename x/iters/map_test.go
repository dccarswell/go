// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestMap(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})
	double := func(v int) int { return v * 2 }

	got := collect(Map(MapFunc[int, int](double))(seq))
	assertIntSlice(t, got, []int{2, 4, 6})
}

func TestMapMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	got := collect(seq.Map(func(v int) int { return v * 2 }))
	assertIntSlice(t, got, []int{2, 4, 6})
}

func TestMapChangesType(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	got := collect(seq.Map(func(v int) string {
		switch v {
		case 1:
			return "one"
		case 2:
			return "two"
		default:
			return "three"
		}
	}))

	want := []string{"one", "two", "three"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestMapEmpty(t *testing.T) {
	got := collect(sliceSeq[int](nil).Map(func(v int) int { return v }))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestMapEarlyTermination(t *testing.T) {
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
	for v := range source.Map(func(v int) int { return v }) {
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

func TestMap2(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})
	swap := func(k string, v int) (int, string) { return v, k }

	got := Map2(MapFunc2[string, int, int, string](swap))(seq)
	keys, vals := collect2(got)

	if len(keys) != 2 || keys[0] != 1 || keys[1] != 2 || vals[0] != "a" || vals[1] != "b" {
		t.Errorf("keys=%v vals=%v, want keys=[1 2] vals=[a b]", keys, vals)
	}
}

func TestMap2Method(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	keys, vals := collect2(seq.Map2(func(k string, v int) (int, string) { return v, k }))

	if len(keys) != 2 || keys[0] != 1 || keys[1] != 2 || vals[0] != "a" || vals[1] != "b" {
		t.Errorf("keys=%v vals=%v, want keys=[1 2] vals=[a b]", keys, vals)
	}
}

func TestMap2Empty(t *testing.T) {
	keys, vals := collect2(sliceSeq2[string, int](nil, nil).Map2(func(k string, v int) (int, string) { return v, k }))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty", keys, vals)
	}
}

func TestMap2EarlyTermination(t *testing.T) {
	calls := 0
	source := Seq2[int, int](func(yield func(int, int) bool) {
		for i := 1; i <= 100; i++ {
			calls++
			if !yield(i, i) {
				return
			}
		}
	})

	count := 0
	for range source.Map2(func(k, v int) (int, int) { return k, v }) {
		count++
		if count == 3 {
			break
		}
	}

	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (early termination not propagated)", calls)
	}
}
