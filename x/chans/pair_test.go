// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package chans

import "testing"

func TestPairBuffered(t *testing.T) {
	send, recv := Pair[int](3)

	send <- 1
	send <- 2
	send <- 3

	for _, want := range []int{1, 2, 3} {
		if got := <-recv; got != want {
			t.Errorf("<-recv = %d, want %d", got, want)
		}
	}
}

func TestPairUnbuffered(t *testing.T) {
	send, recv := Pair[int](0)

	go func() { send <- 42 }()

	if got := <-recv; got != 42 {
		t.Errorf("<-recv = %d, want 42", got)
	}
}

func TestPairSharesUnderlyingChannel(t *testing.T) {
	send, recv := Pair[int](1)

	close(send)

	if v, ok := <-recv; ok || v != 0 {
		t.Errorf("<-recv = (%d, %v), want (0, false) after closing send side", v, ok)
	}
}
