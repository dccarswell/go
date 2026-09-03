package iters

type (
	// ReduceFunc combines an accumulator and the next element into a new
	// accumulator, for [Reduce].
	ReduceFunc[T any] func(T, T) T

	// ReduceFunc2 combines an accumulator and the next (key, value) pair
	// into a new accumulator, for [Reduce2].
	ReduceFunc2[K any, V any] func(V, K, V) V
)

// Seq

// Reduce returns a function that folds a Seq[T] into a single T using f. The
// first element becomes the initial accumulator without calling f; every
// subsequent element is combined via res = f(res, element). If seq is
// empty, Reduce returns the zero value of T without calling f at all — there
// is no error for an empty input.
func Reduce[T any](f ReduceFunc[T]) func(Seq[T]) T {
	return func(seq Seq[T]) T {
		var res T
		first := true
		for i := range seq {
			switch {
			case first:
				res = i
				first = false
			default:
				res = f(res, i)
			}
		}
		return res
	}
}

// Reduce folds s into a single T using f, per the same rules as [Reduce]. It
// is equivalent to Reduce(f)(s).
func (s Seq[T]) Reduce(f ReduceFunc[T]) T {
	return Reduce(f)(s)
}

// Seq2

// Reduce2 returns a function that folds a Seq2[K, V] into a single V using
// f. The first pair's value becomes the initial accumulator without calling
// f; every subsequent pair is combined via res = f(res, key, value). If seq
// is empty, Reduce2 returns the zero value of V without calling f at all —
// there is no error for an empty input.
func Reduce2[K any, V any](f ReduceFunc2[K, V]) func(Seq2[K, V]) V {
	return func(seq Seq2[K, V]) V {
		var res V
		first := true
		for k, v := range seq {
			switch {
			case first:
				res = v
				first = false
			default:
				res = f(res, k, v)
			}
		}
		return res
	}
}

// Reduce2 folds s into a single V using f, per the same rules as [Reduce2]
// the function. It is equivalent to Reduce2(f)(s).
//
// Note the method is named Reduce2, not Reduce — unlike [Seq2.Filter] and
// [Seq2.ForEach], which drop the "2" suffix from their Seq2 method names.
// This is an existing naming inconsistency in the package, not an oversight
// in this doc comment.
func (s Seq2[K, V]) Reduce2(f ReduceFunc2[K, V]) V {
	return Reduce2(f)(s)
}
