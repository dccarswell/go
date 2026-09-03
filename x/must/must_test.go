// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package must

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func TestMustReturnsValueOnNilError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Must panicked: %v", r)
		}
	}()
	if got := Must(42, nil); got != 42 {
		t.Errorf("Must(42, nil) = %d, want 42", got)
	}
}

func TestMustReturnsZeroValueOnNilError(t *testing.T) {
	if got := Must("", nil); got != "" {
		t.Errorf("Must(\"\", nil) = %q, want %q", got, "")
	}
}

func TestMustPreservesPointerIdentity(t *testing.T) {
	x := 7
	p := &x
	got := Must(p, nil)
	if got != p {
		t.Errorf("Must(p, nil) returned a different pointer than p")
	}
}

func TestMustPreservesStructValue(t *testing.T) {
	type point struct{ X, Y int }
	want := point{X: 1, Y: 2}
	if got := Must(want, nil); got != want {
		t.Errorf("Must(%v, nil) = %v, want %v", want, got, want)
	}
}

func TestMustPanicsOnError(t *testing.T) {
	want := errors.New("boom")

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Must did not panic")
		}
		got, ok := r.(error)
		if !ok {
			t.Fatalf("panic value is not an error: %#v", r)
		}
		if !errors.Is(got, want) {
			t.Fatalf("panic value = %v, does not wrap %v", got, want)
		}
		if !strings.Contains(got.Error(), "must caught an error") {
			t.Errorf("panic message = %q, want it to mention %q", got.Error(), "must caught an error")
		}
	}()
	Must(0, want)
}

func TestMustWithStrconvAtoi(t *testing.T) {
	if got := Must(strconv.Atoi("42")); got != 42 {
		t.Errorf("Must(strconv.Atoi(%q)) = %d, want 42", "42", got)
	}
}

func TestMustWithStrconvAtoiInvalidInputPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Must(strconv.Atoi(invalid)) did not panic")
		}
	}()
	Must(strconv.Atoi("not-a-number"))
}

func TestMust2ReturnsValuesOnNilError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Must2 panicked: %v", r)
		}
	}()
	v1, v2 := Must2("a", 1, nil)
	if v1 != "a" || v2 != 1 {
		t.Errorf("Must2(\"a\", 1, nil) = (%q, %d), want (%q, %d)", v1, v2, "a", 1)
	}
}

func TestMust2PanicsOnError(t *testing.T) {
	want := errors.New("boom")

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Must2 did not panic")
		}
		got, ok := r.(error)
		if !ok {
			t.Fatalf("panic value is not an error: %#v", r)
		}
		if !errors.Is(got, want) {
			t.Fatalf("panic value = %v, does not wrap %v", got, want)
		}
		if !strings.Contains(got.Error(), "must2 caught an error") {
			t.Errorf("panic message = %q, want it to mention %q", got.Error(), "must2 caught an error")
		}
	}()
	Must2("", 0, want)
}

func TestMust2WithDifferentTypes(t *testing.T) {
	type point struct{ X, Y int }
	v1, v2 := Must2(point{1, 2}, []string{"a", "b"}, nil)
	if v1 != (point{1, 2}) {
		t.Errorf("v1 = %v, want %v", v1, point{1, 2})
	}
	if len(v2) != 2 || v2[0] != "a" || v2[1] != "b" {
		t.Errorf("v2 = %v, want [a b]", v2)
	}
}
