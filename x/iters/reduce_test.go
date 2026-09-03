// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestReduce(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4})
	sum := func(acc, v int) int { return acc + v }

	got := Reduce(ReduceFunc[int](sum))(seq)
	if got != 10 {
		t.Errorf("got %d, want 10", got)
	}
}

func TestReduceMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4})

	got := seq.Reduce(func(acc, v int) int { return acc + v })
	if got != 10 {
		t.Errorf("got %d, want 10", got)
	}
}

func TestReduceSingleElement(t *testing.T) {
	calls := 0
	seq := sliceSeq([]int{42})

	got := seq.Reduce(func(acc, v int) int {
		calls++
		return acc + v
	})

	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}
	if calls != 0 {
		t.Errorf("reducer called %d times for single-element sequence, want 0", calls)
	}
}

func TestReduceEmpty(t *testing.T) {
	got := sliceSeq[int](nil).Reduce(func(acc, v int) int { return acc + v })
	if got != 0 {
		t.Errorf("got %d, want zero value 0", got)
	}
}

func TestReduce2(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c"}, []int{1, 2, 3})
	sumVals := func(acc int, k string, v int) int { return acc + v }

	got := Reduce2(ReduceFunc2[string, int](sumVals))(seq)
	if got != 6 {
		t.Errorf("got %d, want 6", got)
	}
}

func TestReduce2Method(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c"}, []int{1, 2, 3})

	got := seq.Reduce2(func(acc int, k string, v int) int { return acc + v })
	if got != 6 {
		t.Errorf("got %d, want 6", got)
	}
}

func TestReduce2SingleElement(t *testing.T) {
	calls := 0
	seq := sliceSeq2([]string{"only"}, []int{42})

	got := seq.Reduce2(func(acc int, k string, v int) int {
		calls++
		return acc + v
	})

	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}
	if calls != 0 {
		t.Errorf("reducer called %d times for single-element sequence, want 0", calls)
	}
}

func TestReduce2Empty(t *testing.T) {
	got := sliceSeq2[string, int](nil, nil).Reduce2(func(acc int, k string, v int) int { return acc + v })
	if got != 0 {
		t.Errorf("got %d, want zero value 0", got)
	}
}
