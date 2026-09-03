// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package generics

import "testing"

func TestDefaultWithNonNilPointer(t *testing.T) {
	v := 42
	got := Default(&v, 0)
	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}
}

func TestDefaultWithNilPointer(t *testing.T) {
	var p *int
	got := Default(p, 7)
	if got != 7 {
		t.Errorf("got %d, want 7", got)
	}
}

func TestDefaultString(t *testing.T) {
	var p *string
	got := Default(p, "fallback")
	if got != "fallback" {
		t.Errorf("got %q, want %q", got, "fallback")
	}
}

func TestDerefWithNonNilPointer(t *testing.T) {
	v := 42
	got := Deref(&v)
	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}
}

func TestDerefWithNilPointer(t *testing.T) {
	var p *int
	got := Deref(p)
	if got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}

func TestDerefWithNilPointerStruct(t *testing.T) {
	type point struct{ X, Y int }
	var p *point
	got := Deref(p)
	if got != (point{}) {
		t.Errorf("got %+v, want zero value", got)
	}
}

func TestDerefMatchesDefaultWithZeroValue(t *testing.T) {
	var p *int
	if Deref(p) != Default(p, Zero[int]()) {
		t.Errorf("Deref(p) = %d, want Default(p, Zero[int]()) = %d", Deref(p), Default(p, Zero[int]()))
	}
}
