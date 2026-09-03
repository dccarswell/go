// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestReverse(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4, 5})

	got := collect(Reverse(seq))
	assertIntSlice(t, got, []int{5, 4, 3, 2, 1})
}

func TestReverseEmpty(t *testing.T) {
	got := collect(Reverse(sliceSeq[int](nil)))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestReverseSingleElement(t *testing.T) {
	got := collect(Reverse(sliceSeq([]int{42})))
	assertIntSlice(t, got, []int{42})
}

func TestReverseTwice(t *testing.T) {
	want := []int{1, 2, 3, 4}
	got := collect(Reverse(Reverse(sliceSeq(want))))
	assertIntSlice(t, got, want)
}

// Reverse is documented as not lazy: it drains its source at the point it's
// called, not when the result is later ranged over. This test pins that
// down by consuming only the source up front and confirming the source has
// already been fully read before the returned Seq is ever ranged over.
func TestReverseDrainsSourceEagerly(t *testing.T) {
	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 5; i++ {
			calls++
			if !yield(i) {
				return
			}
		}
	})

	reversed := Reverse(source)

	if calls != 5 {
		t.Errorf("source yielded %d times before the returned Seq was ranged over, want 5 (Reverse should drain eagerly)", calls)
	}

	got := collect(reversed)
	assertIntSlice(t, got, []int{5, 4, 3, 2, 1})
}
