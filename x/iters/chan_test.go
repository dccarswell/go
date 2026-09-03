// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func drainChan[T any](c <-chan T) []T {
	var got []T
	for v := range c {
		got = append(got, v)
	}
	return got
}

func TestToChan(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4})

	got := drainChan(ToChan(seq, 0))

	assertIntSlice(t, got, []int{1, 2, 3, 4})
}

func TestToChanMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4})

	got := drainChan(seq.ToChan(0))

	assertIntSlice(t, got, []int{1, 2, 3, 4})
}

func TestToChanBuffered(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4, 5})

	got := drainChan(ToChan(seq, 10))

	assertIntSlice(t, got, []int{1, 2, 3, 4, 5})
}

func TestToChanEmpty(t *testing.T) {
	got := drainChan(ToChan(sliceSeq[int](nil), 0))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestToChanCloses(t *testing.T) {
	c := ToChan(sliceSeq([]int{1}), 0)

	<-c

	if _, ok := <-c; ok {
		t.Error("channel not closed after all values received")
	}
}
