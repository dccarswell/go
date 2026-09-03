package try

import (
	"errors"
	"testing"
)

func TestTry_NoPanicReturnsNil(t *testing.T) {
	err := Try(func() {})
	if err != nil {
		t.Fatalf("Try() = %v, want nil", err)
	}
}

func TestTry_PanicWithErrorIsWrapped(t *testing.T) {
	want := errors.New("boom")

	err := Try(func() { panic(want) })
	if err == nil {
		t.Fatal("Try() = nil, want an error")
	}

	var rpe RecoveredPanicError
	if !errors.As(err, &rpe) {
		t.Fatalf("Try() error is not a RecoveredPanicError: %#v", err)
	}
	if !errors.Is(rpe.Err, want) {
		t.Fatalf("RecoveredPanicError.Err = %v, want %v", rpe.Err, want)
	}
	if !errors.Is(err, want) {
		t.Fatalf("errors.Is(err, want) = false, want true via Unwrap")
	}
	if got, wantMsg := err.Error(), "wrapped panic: boom"; got != wantMsg {
		t.Fatalf("Error() = %q, want %q", got, wantMsg)
	}
}

func TestTry_PanicWithNonErrorValueIsWrapped(t *testing.T) {
	err := Try(func() { panic("boom") })
	if err == nil {
		t.Fatal("Try() = nil, want an error")
	}

	var rpv RecoveredPanicValue
	if !errors.As(err, &rpv) {
		t.Fatalf("Try() error is not a RecoveredPanicValue: %#v", err)
	}
	if rpv.Value != "boom" {
		t.Fatalf("RecoveredPanicValue.Value = %v, want %q", rpv.Value, "boom")
	}
	if got, want := err.Error(), "panic: boom"; got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
}

func TestTry_PanicWithNonErrorStructValue(t *testing.T) {
	type payload struct{ Code int }

	err := Try(func() { panic(payload{Code: 42}) })
	if err == nil {
		t.Fatal("Try() = nil, want an error")
	}

	var rpv RecoveredPanicValue
	if !errors.As(err, &rpv) {
		t.Fatalf("Try() error is not a RecoveredPanicValue: %#v", err)
	}
	if p, ok := rpv.Value.(payload); !ok || p.Code != 42 {
		t.Fatalf("RecoveredPanicValue.Value = %#v, want payload{Code: 42}", rpv.Value)
	}
}

func TestTry_RecoveryDoesNotLeakBetweenCalls(t *testing.T) {
	if err := Try(func() { panic("first") }); err == nil {
		t.Fatal("first Try() = nil, want an error")
	}
	if err := Try(func() {}); err != nil {
		t.Fatalf("second Try() = %v, want nil", err)
	}
}

func TestRecoveredPanicError_Unwrap(t *testing.T) {
	inner := errors.New("inner")
	e := RecoveredPanicError{Err: inner}
	if got := errors.Unwrap(e); got != inner {
		t.Fatalf("Unwrap() = %v, want %v", got, inner)
	}
}
