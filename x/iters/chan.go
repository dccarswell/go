package iters

// ToChan drains seq onto a new channel from a background goroutine, closing
// the channel once seq is exhausted. buflen sets the channel's buffer size
// (0 for unbuffered). The goroutine sends are paced by the channel: if the
// caller stops reading before the channel is closed, the goroutine blocks
// on its next send indefinitely.
func ToChan[T any](seq Seq[T], buflen int) <-chan T {
	c := make(chan T, buflen)
	go func(s Seq[T]) {
		defer close(c)
		for v := range seq {
			c <- v
		}
	}(seq)
	return c
}

// ToChan drains s onto a new channel, per the same rules as [ToChan]. It is
// equivalent to ToChan(s, buflen).
func (s Seq[T]) ToChan(buflen int) <-chan T {
	return ToChan(s, buflen)
}
