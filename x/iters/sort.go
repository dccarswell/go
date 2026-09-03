package iters

import (
	"cmp"
	"slices"
)

type (
	// SortFunc compares two elements for [Sort], returning a negative number
	// if a should sort before b, zero if they're equal, and a positive
	// number if a should sort after b — the same contract as slices.SortFunc.
	SortFunc[T cmp.Ordered] func(T, T) int
)

// DefaultSortFunc is a [SortFunc] that sorts in ascending order using the
// natural ordering of T, via cmp.Compare. Pass it to [Sort] for the common
// case of a plain ascending sort.
func DefaultSortFunc[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

// Sort returns a function that sorts a Seq[T] according to f (see
// [DefaultSortFunc] for the common ascending case) and returns the result
// as a new Seq[T].
//
// Like [Reverse], Sort is not lazy: sorting requires seeing every element
// first, so Sort fully drains its source (via [Seq.Collect]) at the point
// it's called, not when the result is later ranged over.
func Sort[T cmp.Ordered](f SortFunc[T]) func(Seq[T]) Seq[T] {
	return func(seq Seq[T]) Seq[T] {
		t := seq.Collect()
		slices.SortFunc(t, f)
		return Values(t)
	}
}
