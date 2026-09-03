package generics

import "golang.org/x/exp/constraints"

// Signed reports whether I is a signed integer type. It works by comparing
// the bitwise complement of zero (all bits set) against zero: for a signed
// type this is -1, which is less than zero; for an unsigned type it's the
// type's maximum value, which is not.
func Signed[I constraints.Integer]() bool {
	return ^Zero[I]() < Zero[I]()
}
