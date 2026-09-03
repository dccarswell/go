package iters

import (
	"iter"
	"slices"
)

// Collect fully drains s into a new []T, in iteration order. It is a thin
// wrapper around the stdlib's slices.Collect.
func (s Seq[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](s))
}
