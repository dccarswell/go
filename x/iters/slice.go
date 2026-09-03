package iters

// Values returns a Seq[T] that yields the elements of s in order.
func Values[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

// Unslice returns a Seq2[int, V] that yields the (index, value) pairs of s
// in order, as if ranging over s directly.
func Unslice[V any](s []V) Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

// Unmap returns a Seq2[K, V] that yields every (key, value) pair of m, in
// the unspecified order Go's own map iteration uses.
func Unmap[K comparable, V any](m map[K]V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}
