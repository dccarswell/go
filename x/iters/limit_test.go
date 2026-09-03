// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestLimit(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4, 5})

	got := collect(Limit[int](3)(seq))
	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestLimitMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4, 5})

	got := collect(seq.Limit(3))
	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestLimitZero(t *testing.T) {
	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			calls++
			if !yield(i) {
				return
			}
		}
	})

	got := collect(Limit[int](0)(source))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
	if calls != 0 {
		t.Errorf("source was pulled %d times, want 0 (lim<=0 should not pull from the source at all)", calls)
	}
}

func TestLimitNegative(t *testing.T) {
	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			calls++
			if !yield(i) {
				return
			}
		}
	})

	got := collect(Limit[int](-1)(source))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
	if calls != 0 {
		t.Errorf("source was pulled %d times, want 0 (lim<=0 should not pull from the source at all)", calls)
	}
}

func TestLimitGreaterThanLength(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	got := collect(Limit[int](10)(seq))
	assertIntSlice(t, got, []int{1, 2, 3})
}

// Limit checks the limit again immediately after yielding the lim-th
// element, so it never needs to pull a further element from source just to
// discover the limit has been reached (see Limit's doc comment).
func TestLimitDoesNotOverPullFromSource(t *testing.T) {
	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 100; i++ {
			calls++
			if !yield(i) {
				return
			}
		}
	})

	got := collect(Limit[int](3)(source))

	assertIntSlice(t, got, []int{1, 2, 3})
	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (Limit should not pull beyond the limit)", calls)
	}
}

func TestLimitEarlyTermination(t *testing.T) {
	seq := Limit[int](10)(sliceSeq([]int{1, 2, 3, 4, 5}))

	var got []int
	seq(func(v int) bool {
		got = append(got, v)
		return len(got) < 2
	})

	assertIntSlice(t, got, []int{1, 2})
}

func TestLimit2(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4})

	keys, vals := collect2(Limit2[string, int](2)(seq))
	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" || vals[0] != 1 || vals[1] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[a b] vals=[1 2]", keys, vals)
	}
}

func TestLimit2Method(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4})

	keys, vals := collect2(seq.Limit(2))
	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" || vals[0] != 1 || vals[1] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[a b] vals=[1 2]", keys, vals)
	}
}

func TestLimit2Zero(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	keys, vals := collect2(Limit2[string, int](0)(seq))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty", keys, vals)
	}
}

func TestLimit2GreaterThanLength(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	keys, vals := collect2(Limit2[string, int](10)(seq))
	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" || vals[0] != 1 || vals[1] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[a b] vals=[1 2]", keys, vals)
	}
}

// Same no-over-pull behavior as TestLimitDoesNotOverPullFromSource, for
// Limit2.
func TestLimit2DoesNotOverPullFromSource(t *testing.T) {
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
	for range Limit2[int, int](3)(source) {
		count++
	}

	if count != 3 {
		t.Errorf("yielded %d pairs downstream, want 3", count)
	}
	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (Limit2 should not pull beyond the limit)", calls)
	}
}
