package chans

import "iter"

// Iter2Chan drains seq onto a new unbuffered channel from a background
// goroutine, closing the channel once seq is exhausted. The goroutine sends
// are paced by the channel: if the caller stops reading before the channel
// is closed, the goroutine blocks on its next send indefinitely.
func Iter2Chan[T any](seq iter.Seq[T]) <-chan T {
	c := make(chan T)
	go func(s iter.Seq[T]) {
		defer close(c)
		for v := range s {
			c <- v
		}
	}(seq)
	return c
}

// Chan2Iter returns an iter.Seq[T] that yields the values received from c,
// in the order they arrive, until c is closed. If the consumer stops
// ranging over the returned sequence early, no further values are received
// from c.
func Chan2Iter[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}
