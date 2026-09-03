package must

import "fmt"

func Check(err error) {
	switch err {
	case nil:
	default:
		panic(fmt.Errorf("check caught an error: %w", err))
	}
}
