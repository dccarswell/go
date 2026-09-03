package generics

// Copyright 2026 David Carswell

// Zero returns the zero value of T. It is useful in generic code where a
// zero value is needed as an expression (e.g. a function argument or return
// value) rather than via a "var zero T" declaration.
func Zero[T any]() T {
	var zero T
	return zero
}
