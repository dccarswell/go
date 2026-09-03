package try

import "fmt"

type (
	RecoveredPanicValue struct {
		Value any // The original panic value
	}
	RecoveredPanicError struct {
		Err error // The original error that was panicked with
	}
)

func (e RecoveredPanicValue) Error() string {
	return fmt.Sprintf("panic: %v", e.Value)
}

func (e RecoveredPanicError) Error() string {
	return fmt.Sprintf("wrapped panic: %v", e.Err)
}
func (e RecoveredPanicError) Unwrap() error {
	return e.Err
}
