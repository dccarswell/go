package iters

// Seq

// Limit returns a function that lazily truncates a Seq[T] to at most lim
// elements. If lim is zero or negative, the returned Seq yields nothing —
// seq is not pulled from at all. Limiting is lazy: no work happens until
// the returned Seq is ranged over, and — like [Filter] and [Map] — a
// consumer stopping early stops the source seq too.
//
// Limit never pulls more than lim elements from seq: it checks the limit
// again immediately after yielding the lim-th element, so it does not need
// to pull a further element from seq just to discover the limit has been
// reached.
func Limit[T any](lim int) func(Seq[T]) Seq[T] {
	return func(seq Seq[T]) Seq[T] {
		return func(yield func(T) bool) {
			if lim <= 0 {
				return
			}
			cnt := 0
			for v := range seq {
				cnt++
				if !yield(v) {
					return
				}
				if cnt >= lim {
					return
				}
			}
		}
	}
}

// Limit returns a Seq[T] containing at most the first lim elements of s.
// It is equivalent to Limit[T](lim)(s).
func (s Seq[T]) Limit(lim int) Seq[T] {
	return Limit[T](lim)(s)
}

// Seq2

// Limit2 returns a function that lazily truncates a Seq2[K, V] to at most
// lim (key, value) pairs. If lim is zero or negative, the returned Seq2
// yields nothing — seq is not pulled from at all. Limiting is lazy: no work
// happens until the returned Seq2 is ranged over, and — like [Filter2] and
// [Map2] — a consumer stopping early stops the source seq too.
//
// Limit2 never pulls more than lim pairs from seq, for the same reason
// described in [Limit]'s doc comment.
func Limit2[K any, V any](lim int) func(Seq2[K, V]) Seq2[K, V] {
	return func(seq Seq2[K, V]) Seq2[K, V] {
		return func(yield func(K, V) bool) {
			if lim <= 0 {
				return
			}
			cnt := 0
			for k, v := range seq {
				cnt++
				if !yield(k, v) {
					return
				}
				if cnt >= lim {
					return
				}
			}
		}
	}
}

// Limit returns a Seq2[K, V] containing at most the first lim (key, value)
// pairs of s. It is equivalent to Limit2[K, V](lim)(s).
func (s Seq2[K, V]) Limit(lim int) Seq2[K, V] {
	return Limit2[K, V](lim)(s)
}
