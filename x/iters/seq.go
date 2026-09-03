package iters

import "iter"

type (
	// Seq is a single-value iterator over sequences of T, with the same
	// underlying function type as [iter.Seq]. It is defined locally so this
	// package can attach methods to it; converting between Seq[T] and
	// iter.Seq[T] requires an explicit conversion (see the package doc comment).
	Seq[T any] iter.Seq[T]

	// Seq2 is a key/value iterator over sequences of (K, V) pairs, with the
	// same underlying function type as [iter.Seq2]. It is defined locally so
	// this package can attach methods to it; converting between Seq2[K, V]
	// and iter.Seq2[K, V] requires an explicit conversion (see the package
	// doc comment).
	Seq2[K, V any] iter.Seq2[K, V]
)

// Seq

// Pull converts a Seq into a pull-style iterator: repeated calls to next
// return successive values from seq, with ok false once the sequence is
// exhausted. stop must be called (typically via defer) once the caller is
// done pulling, even if the sequence was not fully consumed. Pull is a thin
// wrapper around the stdlib's iter.Pull.
func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func()) {
	return iter.Pull(iter.Seq[V](seq))
}

//Seq2

// Pull2 converts a Seq2 into a pull-style iterator: repeated calls to next
// return successive (key, value) pairs from seq, with ok false once the
// sequence is exhausted. stop must be called (typically via defer) once the
// caller is done pulling, even if the sequence was not fully consumed. Pull2
// is a thin wrapper around the stdlib's iter.Pull2.
func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func()) {
	return iter.Pull2(iter.Seq2[K, V](seq))
}
