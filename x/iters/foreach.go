package iters

type (
	// ForEachFunc is invoked once per element by [ForEach], for side effects.
	ForEachFunc[T any] func(T)

	// ForEachFunc2 is invoked once per (key, value) pair by [ForEach2], for
	// side effects.
	ForEachFunc2[K any, V any] func(K, V)
)

// Seq

// ForEach returns a function that fully consumes a Seq[T], calling f once
// for each element in order. Unlike [Filter] or [Map], ForEach has no way to
// stop early — it always drains seq completely.
func ForEach[T any](f ForEachFunc[T]) func(Seq[T]) {
	return func(seq Seq[T]) {
		for i := range seq {
			f(i)
		}
	}
}

// ForEach calls f once for each element of s, in order. It is equivalent to
// ForEach(f)(s).
func (s Seq[T]) ForEach(f ForEachFunc[T]) {
	ForEach(f)(s)
}

// Seq2

// ForEach2 returns a function that fully consumes a Seq2[K, V], calling f
// once for each (key, value) pair in order. Unlike [Filter2] or [Map2],
// ForEach2 has no way to stop early — it always drains seq completely.
func ForEach2[K any, V any](f ForEachFunc2[K, V]) func(Seq2[K, V]) {
	return func(seq Seq2[K, V]) {
		for k, v := range seq {
			f(k, v)
		}
	}
}

// ForEach calls f once for each (key, value) pair of s, in order. It is
// equivalent to ForEach2(f)(s).
func (s Seq2[K, V]) ForEach(f ForEachFunc2[K, V]) {
	ForEach2(f)(s)
}
