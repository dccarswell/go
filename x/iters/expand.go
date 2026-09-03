package iters

type (
	// ExpandFunc maps a single element to a sub-sequence for [Expand] to
	// flatten into the result.
	ExpandFunc[T any] func(T) Seq[T]

	// ExpandFunc2 maps a single (key, value) pair to a sub-sequence for
	// [Expand2] to flatten into the result.
	ExpandFunc2[K any, V any] func(K, V) Seq2[K, V]
)

// Seq

// Expand returns a function that lazily maps every element of a Seq[T] to
// a sub-sequence via f, and flattens those sub-sequences into a single
// Seq[T] (also known as flat-map). Each sub-sequence is fully consumed
// before f is called on the next source element, so sub-sequences appear
// in the result in source order, each one contiguous.
//
// Expansion is lazy: no work happens until the returned Seq is ranged
// over. Stopping consumption early — whether while still inside a
// sub-sequence or between them — stops both the current sub-sequence and
// the outer source seq.
func Expand[T any](f ExpandFunc[T]) func(Seq[T]) Seq[T] {
	return func(seq Seq[T]) Seq[T] {
		return func(yield func(T) bool) {
			for v := range seq {
				for v2 := range f(v) {
					if !yield(v2) {
						return
					}
				}
			}
		}
	}
}

// Expand returns a Seq[T] with every element of s expanded and flattened
// via f. It is equivalent to Expand(f)(s).
func (s Seq[T]) Expand(f ExpandFunc[T]) Seq[T] {
	return Expand(f)(s)
}

// Seq2

// Expand2 returns a function that lazily maps every (key, value) pair of a
// Seq2[K, V] to a sub-sequence via f, and flattens those sub-sequences into
// a single Seq2[K, V] (also known as flat-map). Each sub-sequence is fully
// consumed before f is called on the next source pair, so sub-sequences
// appear in the result in source order, each one contiguous.
//
// Expansion is lazy: no work happens until the returned Seq2 is ranged
// over. Stopping consumption early — whether while still inside a
// sub-sequence or between them — stops both the current sub-sequence and
// the outer source seq.
func Expand2[K any, V any](f ExpandFunc2[K, V]) func(Seq2[K, V]) Seq2[K, V] {
	return func(seq Seq2[K, V]) Seq2[K, V] {
		return func(yield func(K, V) bool) {
			for k, v := range seq {
				for k2, v2 := range f(k, v) {
					if !yield(k2, v2) {
						return
					}
				}
			}
		}
	}
}

// Expand returns a Seq2[K, V] with every (key, value) pair of s expanded
// and flattened via f. It is equivalent to Expand2(f)(s).
func (s Seq2[K, V]) Expand(f ExpandFunc2[K, V]) Seq2[K, V] {
	return Expand2(f)(s)
}
