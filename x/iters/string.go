package iters

import (
	"fmt"
	"strings"
)

type (
	// String constrains a type parameter to string and named types whose
	// underlying type is string. It is currently unused by this file's
	// functions (ToString and Stringify are constrained by T any, not
	// String), but is exported for callers who want the same constraint.
	String interface {
		~string
	}
)

// ToString formats a as a string via fmt.Sprintf("%v", a). Values already
// of a string type, and values implementing fmt.Stringer, format as
// expected; anything else falls back to Go's default %v representation for
// that type.
func ToString[T any](a T) string {
	return fmt.Sprintf("%v", a)
}

// Stringify returns a Seq[string] with every element of seq converted via
// [ToString]. Stringifying is lazy: no work happens until the returned Seq
// is ranged over, and stopping consumption early stops the source seq as
// well.
func Stringify[T any](seq Seq[T]) Seq[string] {
	return func(yield func(string) bool) {
		for v := range seq {
			if !yield(ToString(v)) {
				return
			}
		}
	}
}

// Stringify returns a Seq[string] with every element of s converted via
// [ToString]. It is equivalent to Stringify(s).
func (s Seq[T]) Stringify() Seq[string] {
	return Stringify(s)
}

// Join fully drains s and concatenates its elements with sep between each
// one, like strings.Join over a slice. Joining an empty s returns "".
func Join(s Seq[string], sep string) string {
	var b strings.Builder
	for v := range s {
		if b.Len() > 0 {
			b.WriteString(sep)
		}
		b.WriteString(v)
	}
	return b.String()
}
