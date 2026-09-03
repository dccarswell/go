package must

import "fmt"

func Must[T any](val T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must caught an error: %w", err))
	}
	return val
}
func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(fmt.Errorf("must2 caught an error: %w", err))
	}
	return v1, v2
}
