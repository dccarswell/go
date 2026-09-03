package iters

import "fmt"

type (
	// FilterFunc reports whether an element should be kept by [Filter].
	FilterFunc[T any]  func(T) bool
	FilterFuncE[T any] func(T) (bool, error)

	// FilterFunc2 reports whether a (key, value) pair should be kept by [Filter2].
	FilterFunc2[K, V any]  func(K, V) bool
	FilterFuncE2[K, V any] func(K, V) (bool, error)

	filterFuncI[T any] interface {
		FilterFuncE[T] | FilterFunc[T]
	}
	filterFuncI2[K, V any] interface {
		FilterFuncE2[K, V] | FilterFunc2[K, V]
	}
)

// Seq

// Filter returns a function that lazily filters a Seq[T], keeping only the
// elements for which f reports true. Filtering is lazy: no work happens
// until the returned Seq is ranged over, and stopping consumption early
// stops the source seq as well.
// func Filter[T any](f FilterFunc[T]) func(Seq[T]) Seq[T] {
// 	return func(seq Seq[T]) Seq[T] {
// 		return func(yield func(T) bool) {
// 			for i := range seq {
// 				if f(i) {
// 					if !yield(i) {
// 						return
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// errorSetter is implemented by *[Wrapper] and, through the promotion its
// embedding gives it, by *[Wrapper2]: anything that can record a panic
// recovered on its behalf.
type errorSetter interface {
	setError(error)
}

// isErrorSettable reports whether T is a Wrapper-shaped type: *T
// implements [errorSetter], whether because T is [Wrapper] itself or
// because T embeds it (as [Wrapper2] does).
func isErrorSettable[T any]() bool {
	var zero T
	_, ok := any(&zero).(errorSetter)
	return ok
}

// Filter returns a function that lazily filters a Seq[T], keeping only the
// elements for which f reports true. Filtering is lazy: no work happens
// until the returned Seq is ranged over, and stopping consumption early
// stops the source seq as well.
//
// Whether T is a Wrapper-shaped type (as with [Wrapper] and [Wrapper2]) is
// checked once, outside the loop, and if so, Filter also recovers any
// panic from f: rather than aborting iteration, it records the panic on
// the element's Error field and yields the element, so a single
// misbehaving predicate can't stop the rest of the sequence.
func Filter[T any](f FilterFunc[T]) func(Seq[T]) Seq[T] {
	recoverable := isErrorSettable[T]()
	return func(seq Seq[T]) Seq[T] {
		return func(yield func(T) bool) {
			for i := range seq {
				var keep bool
				if recoverable {
					keep = filterRecovering(f, &i)
				} else {
					keep = f(i)
				}
				if keep {
					if !yield(i) {
						return
					}
				}
			}
		}
	}
}

// filterRecovering calls f(*i), recovering any panic and, rather than
// letting it propagate, recording it on *i's Error field via [errorSetter].
// It always reports true when it recovers a panic, so the now-failed
// element is yielded instead of dropped.
func filterRecovering[T any](f FilterFunc[T], i *T) (keep bool) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			any(i).(errorSetter).setError(err)
			keep = true
		}
	}()
	return f(*i)
}

// Filter returns a Seq[T] containing only the elements of s for which f
// reports true. It is equivalent to Filter(f)(s).
func (s Seq[T]) Filter(f FilterFunc[T]) Seq[T] {
	return Filter(f)(s)
}

// Seq2

// Filter2 returns a function that lazily filters a Seq2[K, V], keeping only
// the (key, value) pairs for which f reports true. Filtering is lazy: no
// work happens until the returned Seq2 is ranged over, and stopping
// consumption early stops the source seq as well.
func Filter2[K, V any](f FilterFunc2[K, V]) func(Seq2[K, V]) Seq2[K, V] {
	return func(seq Seq2[K, V]) Seq2[K, V] {
		return func(yield func(K, V) bool) {
			for k, v := range seq {
				if f(k, v) {
					if !yield(k, v) {
						return
					}
				}
			}
		}
	}
}

// Filter returns a Seq2[K, V] containing only the (key, value) pairs of s
// for which f reports true. It is equivalent to Filter2(f)(s).
func (s Seq2[K, V]) Filter(f FilterFunc2[K, V]) Seq2[K, V] {
	return Filter2(f)(s)
}
