package generics

// Default returns *p, or def if p is nil.
func Default[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// Deref returns *p, or the zero value of T if p is nil. It is equivalent to
// Default(p, Zero[T]()).
func Deref[T any](p *T) T {
	if p == nil {
		return Zero[T]()
	}
	return *p
}
