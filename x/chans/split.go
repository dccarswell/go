package chans

// Split returns send-only and receive-only views of the existing channel c.
// Both views share the same underlying channel: a value sent on the first
// is received on the second, and closing the first is observable as a
// closed channel on the second. Unlike [Pair], Split does not create a
// channel — it only narrows the direction of one you already have,
// including any values already buffered in it.
func Split[T any](c chan T) (chan<- T, <-chan T) {
	return c, c
}
