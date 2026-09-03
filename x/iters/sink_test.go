// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestSink(t *testing.T) {
	count := 0
	seq := Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 5; i++ {
			count++
			if !yield(i) {
				return
			}
		}
	})

	Sink(seq)

	if count != 5 {
		t.Errorf("consumed %d elements, want 5", count)
	}
}

func TestSinkEmpty(t *testing.T) {
	Sink(sliceSeq[int](nil))
}

func TestSink2(t *testing.T) {
	count := 0
	seq := Seq2[int, int](func(yield func(int, int) bool) {
		for i := 1; i <= 5; i++ {
			count++
			if !yield(i, i*10) {
				return
			}
		}
	})

	Sink2(seq)

	if count != 5 {
		t.Errorf("consumed %d elements, want 5", count)
	}
}

func TestSink2Empty(t *testing.T) {
	Sink2(sliceSeq2[int, int](nil, nil))
}
