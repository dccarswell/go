// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package chans

import (
	"iter"
	"testing"
)

func sliceIterSeq[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range items {
			if !yield(v) {
				return
			}
		}
	}
}

func drain[T any](c <-chan T) []T {
	var got []T
	for v := range c {
		got = append(got, v)
	}
	return got
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

func TestIter2Chan(t *testing.T) {
	seq := sliceIterSeq([]int{1, 2, 3, 4})

	got := drain(Iter2Chan(seq))

	assertIntSlice(t, got, []int{1, 2, 3, 4})
}

func TestIter2ChanEmpty(t *testing.T) {
	got := drain(Iter2Chan(sliceIterSeq[int](nil)))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestIter2ChanCloses(t *testing.T) {
	c := Iter2Chan(sliceIterSeq([]int{1}))

	<-c

	if _, ok := <-c; ok {
		t.Error("channel not closed after all values received")
	}
}

func TestChan2Iter(t *testing.T) {
	c := make(chan int, 4)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	var got []int
	for v := range Chan2Iter[int](c) {
		got = append(got, v)
	}

	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestChan2IterEmpty(t *testing.T) {
	c := make(chan int)
	close(c)

	var got []int
	for v := range Chan2Iter[int](c) {
		got = append(got, v)
	}
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestChan2IterEarlyTermination(t *testing.T) {
	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}

	var got []int
	for v := range Chan2Iter[int](c) {
		got = append(got, v)
		if len(got) == 2 {
			break
		}
	}

	assertIntSlice(t, got, []int{1, 2})
	if len(c) != 3 {
		t.Errorf("channel has %d buffered values left, want 3 (early termination drained too much)", len(c))
	}
}

func TestIter2ChanChan2IterRoundTrip(t *testing.T) {
	want := []int{5, 4, 3, 2, 1}
	seq := sliceIterSeq(want)

	var got []int
	for v := range Chan2Iter[int](Iter2Chan(seq)) {
		got = append(got, v)
	}

	assertIntSlice(t, got, want)
}
