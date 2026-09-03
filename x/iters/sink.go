package iters

// Seq

// Sink fully drains seq, discarding every element. It is useful for forcing
// a lazily-built chain of operations (e.g. [Filter] or [Map]) to run for
// their side effects alone.
func Sink[T any](seq Seq[T]) {
	for _ = range seq {
	}
}

// Seq2

// Sink2 fully drains seq, discarding every (key, value) pair. It is useful
// for forcing a lazily-built chain of operations (e.g. [Filter2] or [Map2])
// to run for their side effects alone.
func Sink2[K, V any](seq Seq2[K, V]) {
	for _, _ = range seq {
	}
}
