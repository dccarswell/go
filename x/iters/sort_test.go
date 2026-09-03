// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestSortAscending(t *testing.T) {
	seq := sliceSeq([]int{5, 3, 1, 4, 2})

	got := collect(Sort(DefaultSortFunc[int])(seq))
	assertIntSlice(t, got, []int{1, 2, 3, 4, 5})
}

func TestSortCustomOrder(t *testing.T) {
	seq := sliceSeq([]int{5, 3, 1, 4, 2})
	descending := func(a, b int) int { return b - a }

	got := collect(Sort(SortFunc[int](descending))(seq))
	assertIntSlice(t, got, []int{5, 4, 3, 2, 1})
}

func TestSortEmpty(t *testing.T) {
	got := collect(Sort(DefaultSortFunc[int])(sliceSeq[int](nil)))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestSortSingleElement(t *testing.T) {
	got := collect(Sort(DefaultSortFunc[int])(sliceSeq([]int{42})))
	assertIntSlice(t, got, []int{42})
}

func TestSortAlreadySorted(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := collect(Sort(DefaultSortFunc[int])(sliceSeq(want)))
	assertIntSlice(t, got, want)
}

func TestSortStrings(t *testing.T) {
	seq := sliceSeq([]string{"banana", "apple", "cherry"})

	got := collect(Sort(DefaultSortFunc[string])(seq))

	want := []string{"apple", "banana", "cherry"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

// Sort is documented as not lazy: it drains its source at the point the
// curried function is applied to a Seq, not when the result is later ranged
// over. This test pins that down the same way TestReverseDrainsSourceEagerly
// does for Reverse.
func TestSortDrainsSourceEagerly(t *testing.T) {
	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for _, v := range []int{3, 1, 2} {
			calls++
			if !yield(v) {
				return
			}
		}
	})

	sorted := Sort(DefaultSortFunc[int])(source)

	if calls != 3 {
		t.Errorf("source yielded %d times before the returned Seq was ranged over, want 3 (Sort should drain eagerly)", calls)
	}

	got := collect(sorted)
	assertIntSlice(t, got, []int{1, 2, 3})
}
