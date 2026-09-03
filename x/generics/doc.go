// This file's package doc comment was written by Claude (Anthropic's AI assistant).

// Package generics provides small, general-purpose generic helpers that fill
// gaps in what Go's type system otherwise makes awkward to express inline:
//
//   - [Zero]: the zero value of a type parameter, as a value rather than via
//     a "var zero T" declaration.
//   - [MinMax]: the minimum and maximum representable values of an integer or
//     float type.
//   - [Signed]: whether an integer type is signed.
//   - [Ptr]: the address of a value, including literals and other
//     non-addressable expressions.
//   - [Default] and [Deref]: safely read through a possibly-nil pointer,
//     substituting a fallback value or the zero value respectively.
package generics
