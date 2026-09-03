// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestCount(t *testing.T) {
	got := Count(sliceSeq([]int{1, 2, 3, 4}))
	if got != 4 {
		t.Errorf("got %d, want 4", got)
	}
}

func TestCountMethod(t *testing.T) {
	got := sliceSeq([]int{1, 2, 3, 4}).Count()
	if got != 4 {
		t.Errorf("got %d, want 4", got)
	}
}

func TestCountEmpty(t *testing.T) {
	got := sliceSeq[int](nil).Count()
	if got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}

func TestCount2(t *testing.T) {
	got := Count2(sliceSeq2([]string{"a", "b", "c"}, []int{1, 2, 3}))
	if got != 3 {
		t.Errorf("got %d, want 3", got)
	}
}

func TestCount2Method(t *testing.T) {
	got := sliceSeq2([]string{"a", "b", "c"}, []int{1, 2, 3}).Count()
	if got != 3 {
		t.Errorf("got %d, want 3", got)
	}
}

func TestCount2Empty(t *testing.T) {
	got := sliceSeq2[string, int](nil, nil).Count()
	if got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}
