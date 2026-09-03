package iters

type (
	// MapFunc transforms a T1 into a T2 for [Map].
	MapFunc[T1 any, T2 any] func(T1) T2

	// MapFunc2 transforms a (K1, V1) pair into a (K2, V2) pair for [Map2].
	MapFunc2[K1 any, V1 any, K2 any, V2 any] func(K1, V1) (K2, V2)
)

// Seq

// Map returns a function that lazily transforms every element of a Seq[T1]
// into a Seq[T2] via f. Mapping is lazy: no work happens until the returned
// Seq is ranged over, and stopping consumption early stops the source seq
// as well.
func Map[T1 any, T2 any](f MapFunc[T1, T2]) func(Seq[T1]) Seq[T2] {
	return func(seq Seq[T1]) Seq[T2] {
		return func(yield func(T2) bool) {
			for i := range seq {
				if !yield(f(i)) {
					return
				}
			}
		}
	}
}

// Map returns a Seq[T2] with every element of s transformed by f. It is
// equivalent to Map(f)(s).
func (s Seq[T1]) Map[T2 any](f MapFunc[T1, T2]) Seq[T2] {
	return Map(f)(s)
}

// Seq2

// Map2 returns a function that lazily transforms every (key, value) pair of
// a Seq2[K1, V1] into a Seq2[K2, V2] via f, which may change both the key
// and value types. Mapping is lazy: no work happens until the returned Seq2
// is ranged over, and stopping consumption early stops the source seq as
// well.
func Map2[K1 any, V1 any, K2 any, V2 any](f MapFunc2[K1, V1, K2, V2]) func(Seq2[K1, V1]) Seq2[K2, V2] {
	return func(seq Seq2[K1, V1]) Seq2[K2, V2] {
		return func(yield func(K2, V2) bool) {
			for k, v := range seq {
				k2, v2 := f(k, v)
				if !yield(k2, v2) {
					return
				}
			}
		}
	}
}

// Map2 returns a Seq2[K2, V2] with every (key, value) pair of s transformed
// by f. It is equivalent to Map2(f)(s).
func (s Seq2[K1, V1]) Map2[K2 any, V2 any](f MapFunc2[K1, V1, K2, V2]) Seq2[K2, V2] {
	return Map2(f)(s)
}