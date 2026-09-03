package iters

type (
	// CoalesceFunc combines a (key, value) pair into a single T, for [Coalesce].
	CoalesceFunc[K any, V any, T any] func(K, V) T
)

// Coalesce returns a function that lazily collapses a Seq2[K, V] into a
// Seq[T] by applying f to every (key, value) pair. Coalescing is lazy: no
// work happens until the returned Seq is ranged over, and stopping
// consumption early stops the source seq as well.
func Coalesce[K any, V any, T any](f CoalesceFunc[K, V, T]) func(Seq2[K, V]) Seq[T] {
	return func(seq Seq2[K, V]) Seq[T] {
		return func(yield func(T) bool) {
			for k, v := range seq {
				if !yield(f(k, v)) {
					return
				}
			}
		}
	}
}

// Coalesce returns a Seq[T] with every (key, value) pair of s collapsed
// into a single T via f. It is equivalent to Coalesce(f)(s).
func (s Seq2[K, V]) Coalesce[T any](f CoalesceFunc[K, V, T]) Seq[T] {
	return Coalesce(f)(s)
}
