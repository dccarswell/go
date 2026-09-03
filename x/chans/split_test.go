// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package chans

import "testing"

func TestSplitBuffered(t *testing.T) {
	c := make(chan int, 3)
	send, recv := Split(c)

	send <- 1
	send <- 2
	send <- 3

	for _, want := range []int{1, 2, 3} {
		if got := <-recv; got != want {
			t.Errorf("<-recv = %d, want %d", got, want)
		}
	}
}

func TestSplitUnbuffered(t *testing.T) {
	c := make(chan int)
	send, recv := Split(c)

	go func() { send <- 42 }()

	if got := <-recv; got != 42 {
		t.Errorf("<-recv = %d, want 42", got)
	}
}

func TestSplitSharesUnderlyingChannel(t *testing.T) {
	c := make(chan int, 1)
	send, recv := Split(c)

	close(send)

	if v, ok := <-recv; ok || v != 0 {
		t.Errorf("<-recv = (%d, %v), want (0, false) after closing send side", v, ok)
	}
}

func TestSplitOfExistingChannelContents(t *testing.T) {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	_, recv := Split(c)

	if got := <-recv; got != 1 {
		t.Errorf("<-recv = %d, want 1", got)
	}
	if got := <-recv; got != 2 {
		t.Errorf("<-recv = %d, want 2", got)
	}
}
