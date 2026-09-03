// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestSumInts(t *testing.T) {
	got := sliceSeq([]int{1, 2, 3, 4}).Reduce(Sum)
	if got != 10 {
		t.Errorf("got %d, want 10", got)
	}
}

func TestSumFloats(t *testing.T) {
	got := sliceSeq([]float64{1.5, 2.5, 3.0}).Reduce(Sum)
	if got != 7.0 {
		t.Errorf("got %v, want 7.0", got)
	}
}

func TestSumSingleElement(t *testing.T) {
	got := sliceSeq([]int{42}).Reduce(Sum)
	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}
}

func TestSumEmpty(t *testing.T) {
	got := sliceSeq[int](nil).Reduce(Sum)
	if got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}

func TestSumAsPlainFunction(t *testing.T) {
	got := Sum(3, 4)
	if got != 7 {
		t.Errorf("Sum(3, 4) = %d, want 7", got)
	}
}
