// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package must

import (
	"errors"
	"strings"
	"testing"
)

func TestCheckNilDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Check(nil) panicked: %v", r)
		}
	}()
	Check(nil)
}

func TestCheckErrorPanics(t *testing.T) {
	want := errors.New("boom")

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Check(err) did not panic")
		}
		got, ok := r.(error)
		if !ok {
			t.Fatalf("panic value is not an error: %#v", r)
		}
		if !errors.Is(got, want) {
			t.Fatalf("panic value = %v, does not wrap %v", got, want)
		}
		if !strings.Contains(got.Error(), "check caught an error") {
			t.Errorf("panic message = %q, want it to mention %q", got.Error(), "check caught an error")
		}
	}()
	Check(want)
}
