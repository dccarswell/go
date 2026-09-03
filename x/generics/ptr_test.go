// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package generics

import "testing"

func TestPtrPointsToEqualValue(t *testing.T) {
	p := Ptr(42)
	if p == nil {
		t.Fatal("Ptr(42) returned nil")
	}
	if *p != 42 {
		t.Errorf("*Ptr(42) = %d, want 42", *p)
	}
}

func TestPtrOfLiteral(t *testing.T) {
	// The point of Ptr is exactly that Ptr(5) compiles where &5 would not.
	p := Ptr(5)
	if *p != 5 {
		t.Errorf("*Ptr(5) = %d, want 5", *p)
	}
}

func TestPtrOfExpression(t *testing.T) {
	x := 10
	p := Ptr(x + 1)
	if *p != 11 {
		t.Errorf("*Ptr(x+1) = %d, want 11", *p)
	}
}

func TestPtrIsIndependentCopy(t *testing.T) {
	x := 1
	p := Ptr(x)
	x = 2
	if *p != 1 {
		t.Errorf("*Ptr(x) = %d after mutating x, want 1 (Ptr should copy the value)", *p)
	}
}

func TestPtrString(t *testing.T) {
	p := Ptr("hello")
	if *p != "hello" {
		t.Errorf("*Ptr(%q) = %q, want %q", "hello", *p, "hello")
	}
}
