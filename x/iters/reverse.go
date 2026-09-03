package iters

import "slices"

// Reverse returns a Seq[T] yielding the elements of seq in reverse order.
//
// Unlike the rest of this package's operations, Reverse is not lazy: since
// the last element can't be yielded until every earlier element has been
// seen, Reverse must fully drain seq (via [Seq.Collect]) before it can
// return anything. This happens once, at the point Reverse is called, not
// when the result is later ranged over.
func Reverse[T any](seq Seq[T]) Seq[T] {
	t := seq.Collect()
	slices.Reverse(t)
	return Values(t)
}
