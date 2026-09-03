package iters

type (
	EndFunc[T any]     func(T) bool
	EndFunc2[K, V any] func(K, V) bool
)

func End[T any](f EndFunc[T]) func(Seq[T]) Seq[T] {
	return func(seq Seq[T]) Seq[T] {
		return func(yield func(T) bool) {
			end := false
			for i := range seq {
				if end {
					return
				}
				end = f(i)
				if !end {
					if !yield(i) {
						return
					}
				}
			}
		}
	}
}
func (s Seq[T]) End(f EndFunc[T]) Seq[T] {
	return End(f)(s)
}

func End2[K, V any](f EndFunc2[K, V]) func(Seq2[K, V]) Seq2[K, V] {
	return func(seq Seq2[K, V]) Seq2[K, V] {
		return func(yield func(K, V) bool) {
			end := false
			for k, v := range seq {
				if end {
					return
				}
				end = f(k, v)
				if !end {
					if !yield(k, v) {
						return
					}
				}
			}
		}
	}
}
func (s Seq2[K, V]) End(f EndFunc2[K, V]) Seq2[K, V] {
	return End2(f)(s)
}
