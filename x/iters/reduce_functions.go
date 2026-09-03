package iters

import "golang.org/x/exp/constraints"

// Sum is a [ReduceFunc] that adds each element to the running total. Pass
// it to [Reduce] (e.g. seq.Reduce(iters.Sum)) to total a Seq[N] of any
// integer or float type. Per Reduce's rules, summing an empty sequence
// yields the zero value of N.
func Sum[N constraints.Integer | constraints.Float](res N, el N) N {
	return res + el
}
