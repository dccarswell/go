package chans

// Pair creates a new channel with the given buffer length and returns
// send-only and receive-only views of it. Both views share the same
// underlying channel: a value sent on the first is received on the second,
// and closing the first is observable as a closed channel on the second.
func Pair[T any](buflen int) (chan<- T, <-chan T) {
	c := make(chan T, buflen)
	return c, c
}
