package try

func trap(e *error) {
	r := recover()
	switch r := r.(type) {
	case nil:
	case error:
		*e = RecoveredPanicError{Err: r}
	default:
		*e = RecoveredPanicValue{Value: r}
	}
}
func Try(f TryFunc) (err error) {
	defer trap(&err)
	f()
	return
}
