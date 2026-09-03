// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package generics

import (
	"reflect"
	"testing"
)

func TestZeroNumeric(t *testing.T) {
	if got := Zero[int](); got != 0 {
		t.Errorf("Zero[int]() = %v, want 0", got)
	}
	if got := Zero[int8](); got != 0 {
		t.Errorf("Zero[int8]() = %v, want 0", got)
	}
	if got := Zero[uint64](); got != 0 {
		t.Errorf("Zero[uint64]() = %v, want 0", got)
	}
	if got := Zero[float64](); got != 0 {
		t.Errorf("Zero[float64]() = %v, want 0", got)
	}
	if got := Zero[complex128](); got != 0 {
		t.Errorf("Zero[complex128]() = %v, want 0", got)
	}
}

func TestZeroString(t *testing.T) {
	if got := Zero[string](); got != "" {
		t.Errorf("Zero[string]() = %q, want \"\"", got)
	}
}

func TestZeroBool(t *testing.T) {
	if got := Zero[bool](); got {
		t.Errorf("Zero[bool]() = %v, want false", got)
	}
}

func TestZeroPointer(t *testing.T) {
	if got := Zero[*int](); got != nil {
		t.Errorf("Zero[*int]() = %v, want nil", got)
	}
}

// Reference kinds whose zero value is nil but which are not comparable with !=
// against anything other than nil.
func TestZeroNilable(t *testing.T) {
	if got := Zero[[]int](); got != nil {
		t.Errorf("Zero[[]int]() = %v, want nil", got)
	}
	if got := Zero[map[string]int](); got != nil {
		t.Errorf("Zero[map[string]int]() = %v, want nil", got)
	}
	if got := Zero[chan int](); got != nil {
		t.Errorf("Zero[chan int]() = %v, want nil", got)
	}
	if got := Zero[func()](); got != nil {
		t.Error("Zero[func()]() is non-nil, want nil")
	}
	if got := Zero[error](); got != nil {
		t.Errorf("Zero[error]() = %v, want nil", got)
	}
	if got := Zero[any](); got != nil {
		t.Errorf("Zero[any]() = %v, want nil", got)
	}
}

type point struct {
	X, Y int
	Name string
	Tags []string
}

func TestZeroStruct(t *testing.T) {
	got := Zero[point]()
	want := point{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Zero[point]() = %+v, want %+v", got, want)
	}
	if got.Tags != nil {
		t.Errorf("Zero[point]().Tags = %v, want nil", got.Tags)
	}
}

func TestZeroArray(t *testing.T) {
	got := Zero[[3]int]()
	want := [3]int{}
	if got != want {
		t.Errorf("Zero[[3]int]() = %v, want %v", got, want)
	}
}

type myInt int

func TestZeroNamedType(t *testing.T) {
	if got := Zero[myInt](); got != 0 {
		t.Errorf("Zero[myInt]() = %v, want 0", got)
	}
}

// Zero must return a fresh value each call; mutating the result of one call
// must not affect subsequent calls.
func TestZeroIndependentValues(t *testing.T) {
	a := Zero[point]()
	a.X = 42
	a.Tags = append(a.Tags, "mutated")

	b := Zero[point]()
	if b.X != 0 || b.Tags != nil {
		t.Errorf("Zero[point]() = %+v after mutating a previous result, want zero value", b)
	}
}

// The value returned for any type T must match reflect's notion of that type's
// zero value.
func TestZeroMatchesReflect(t *testing.T) {
	assertReflectZero[int](t)
	assertReflectZero[string](t)
	assertReflectZero[bool](t)
	assertReflectZero[float32](t)
	assertReflectZero[*point](t)
	assertReflectZero[point](t)
	assertReflectZero[[2]string](t)
	assertReflectZero[[]byte](t)
	assertReflectZero[map[int]int](t)
	assertReflectZero[struct{ A int }](t)
}

func assertReflectZero[T any](t *testing.T) {
	t.Helper()
	got := Zero[T]()
	want := reflect.Zero(reflect.TypeFor[T]()).Interface()
	if !reflect.DeepEqual(any(got), want) {
		t.Errorf("Zero[%s]() = %#v, want %#v", reflect.TypeFor[T](), got, want)
	}
}

// Zero should be usable where a value of T is expected without explicit
// instantiation at the call site.
func TestZeroInferredInGenericContext(t *testing.T) {
	if got := firstOrZero([]int{}); got != 0 {
		t.Errorf("firstOrZero([]int{}) = %v, want 0", got)
	}
	if got := firstOrZero([]int{7, 8}); got != 7 {
		t.Errorf("firstOrZero([]int{7, 8}) = %v, want 7", got)
	}
	if got := firstOrZero([]string{}); got != "" {
		t.Errorf("firstOrZero([]string{}) = %q, want \"\"", got)
	}
}

func firstOrZero[T any](s []T) T {
	if len(s) == 0 {
		return Zero[T]()
	}
	return s[0]
}
