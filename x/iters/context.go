package iters

import "context"

// Seq

// Monitor returns a function that lazily wraps a Seq[T] so that iteration
// stops early once ctx is done, in addition to stopping on the downstream
// consumer's own request. ctx is checked before each element is yielded, so
// an already-done ctx stops iteration without producing that element. Like
// [Filter] and [Map], stopping — for either reason — stops the source seq
// too.
//
// Even if ctx is already done before iteration starts, Monitor still pulls
// one element from seq before it gets a chance to check ctx (an unavoidable
// consequence of range-over-func: the first element must be produced before
// any loop body can run), though that element is discarded rather than
// yielded downstream.
func Monitor[T any](ctx context.Context) func(Seq[T]) Seq[T] {
	return func(seq Seq[T]) Seq[T] {
		return func(yield func(T) bool) {
			for i := range seq {
				select {
				case <-ctx.Done():
					return
				default:
					if !yield(i) {
						return
					}
				}
			}
		}
	}
}

// Monitor returns a Seq[T] that stops iterating s once ctx is done. It is
// equivalent to Monitor[T](ctx)(s).
func (s Seq[T]) Monitor(ctx context.Context) Seq[T] {
	return Monitor[T](ctx)(s)
}

// Seq2

// Monitor2 returns a function that lazily wraps a Seq2[K, V] so that
// iteration stops early once ctx is done, in addition to stopping on the
// downstream consumer's own request. ctx is checked before each pair is
// yielded, so an already-done ctx stops iteration without producing that
// pair. Like [Filter2] and [Map2], stopping — for either reason — stops the
// source seq too.
//
// Even if ctx is already done before iteration starts, Monitor2 still pulls
// one pair from seq before it gets a chance to check ctx, for the same
// reason described in [Monitor]'s doc comment.
func Monitor2[K, V any](ctx context.Context) func(Seq2[K, V]) Seq2[K, V] {
	return func(seq Seq2[K, V]) Seq2[K, V] {
		return func(yield func(K, V) bool) {
			for k, v := range seq {
				select {
				case <-ctx.Done():
					return
				default:
					if !yield(k, v) {
						return
					}
				}
			}
		}
	}
}

// Monitor returns a Seq2[K, V] that stops iterating s once ctx is done. It
// is equivalent to Monitor2[K, V](ctx)(s).
func (s Seq2[K, V]) Monitor(ctx context.Context) Seq2[K, V] {
	return Monitor2[K, V](ctx)(s)
}
