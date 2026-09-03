package iters

type (
	wrapper[T any] struct {
		Value T
		Error error
	}
	wrapper2[K any, V any] struct {
		wrapper[V]
		Key K
	}
	iwrapper[T any] interface {
		setError(error)
	}

	WrapFunc[T any]         func(T) wrapper[T]
	WrapFunc2[K any, V any] func(K, V) wrapper2[K, V]
)

func (w *wrapper[T]) setError(err error) {
	w.Error = err
}

func (w *wrapper[T]) Ok() bool {
	return w.Error == nil
}

func Wrap[T any](f WrapFunc[T]) func(Seq[T]) Seq[wrapper[T]] {
	return func(seq Seq[T]) Seq[wrapper[T]] {
		return func(yield func(wrapper[T]) bool) {
			for i := range seq {
				if !yield(f(i)) {
					return
				}
			}
		}
	}
}

func (s Seq[T]) Wrap(f WrapFunc[T]) Seq[wrapper[T]] {
	return func(yield func(wrapper[T]) bool) {
		for i := range s {
			if !yield(T2(f(i))) {
				return
			}
		}
	}
}

func Wrap2[K any, V any](f WrapFunc2[K, V]) func(Seq2[K, V]) Seq[wrapper2[K, V]] {
	return func(seq Seq2[K, V]) Seq[wrapper2[K, V]] {
		return func(yield func(wrapper2[K, V]) bool) {
			for k, v := range seq {
				if !yield(f(k, v)) {
					return
				}
			}
		}
	}
}

// Wrap2 is the method form of [Wrap2]: it lazily applies f to every
// (key, value) pair of s, returning a Seq of the wrapped results. It is
// equivalent to Wrap2(f)(s).
func (s Seq2[K, V]) Wrap2(f WrapFunc2[K, V]) Seq[wrapper2[K, V]] {
	return Wrap2(f)(s)
}
