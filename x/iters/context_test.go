// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import (
	"context"
	"testing"
)

func TestMonitorPassesThroughWhenNotDone(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	got := collect(Monitor[int](context.Background())(seq))
	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestMonitorMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	got := collect(seq.Monitor(context.Background()))
	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestMonitorStopsOnAlreadyDoneContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	seq := sliceSeq([]int{1, 2, 3})

	got := collect(Monitor[int](ctx)(seq))
	if len(got) != 0 {
		t.Errorf("got %v, want empty (context was already done)", got)
	}
}

// Documents the pull-before-check behavior noted in Monitor's doc comment:
// even with an already-done ctx, one element is still pulled from source
// (range-over-func must produce it before the loop body can check ctx) —
// it just never reaches the downstream consumer.
func TestMonitorPullsOnceEvenOnAlreadyDoneContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			calls++
			if !yield(i) {
				return
			}
		}
	})

	got := collect(Monitor[int](ctx)(source))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
	if calls != 1 {
		t.Errorf("source was pulled %d times, want 1", calls)
	}
}

func TestMonitorStopsSourceWhenCancelledMidIteration(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 100; i++ {
			calls++
			if i == 3 {
				cancel()
			}
			if !yield(i) {
				return
			}
		}
	})

	collect(Monitor[int](ctx)(source))

	if calls > 4 {
		t.Errorf("source yielded %d times after cancellation, want at most 4 (cancel on the 3rd element, one more pull to observe it)", calls)
	}
}

func TestMonitorEmpty(t *testing.T) {
	got := collect(Monitor[int](context.Background())(sliceSeq[int](nil)))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestMonitor2PassesThroughWhenNotDone(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	keys, vals := collect2(Monitor2[string, int](context.Background())(seq))
	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" || vals[0] != 1 || vals[1] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[a b] vals=[1 2]", keys, vals)
	}
}

func TestMonitor2Method(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	keys, vals := collect2(seq.Monitor(context.Background()))
	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" || vals[0] != 1 || vals[1] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[a b] vals=[1 2]", keys, vals)
	}
}

func TestMonitor2StopsOnAlreadyDoneContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	keys, vals := collect2(Monitor2[string, int](ctx)(seq))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty (context was already done)", keys, vals)
	}
}

// Same pull-before-check behavior as TestMonitorPullsOnceEvenOnAlreadyDoneContext,
// for Monitor2.
func TestMonitor2PullsOnceEvenOnAlreadyDoneContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	calls := 0
	source := Seq2[int, int](func(yield func(int, int) bool) {
		for i := 1; i <= 3; i++ {
			calls++
			if !yield(i, i) {
				return
			}
		}
	})

	keys, vals := collect2(Monitor2[int, int](ctx)(source))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty", keys, vals)
	}
	if calls != 1 {
		t.Errorf("source was pulled %d times, want 1", calls)
	}
}

func TestMonitor2Empty(t *testing.T) {
	keys, vals := collect2(Monitor2[string, int](context.Background())(sliceSeq2[string, int](nil, nil)))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty", keys, vals)
	}
}
