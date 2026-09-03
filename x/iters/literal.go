package iters

// Seq

// Create returns a Seq[T] that yields the given elements in order. It is
// the variadic-literal counterpart to [Values], for building a sequence
// directly from arguments instead of an existing slice.
func Create[T any](s ...T) Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}
