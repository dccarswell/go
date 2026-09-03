// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package generics

import (
	"math"
	"testing"
)

func TestMinMaxUint8(t *testing.T) {
	min, max := MinMax[uint8]()
	if min != 0 {
		t.Errorf("min = %v, want 0", min)
	}
	if max != math.MaxUint8 {
		t.Errorf("max = %v, want %v", max, uint8(math.MaxUint8))
	}
}

func TestMinMaxInt8(t *testing.T) {
	min, max := MinMax[int8]()
	if min != math.MinInt8 {
		t.Errorf("min = %v, want %v", min, int8(math.MinInt8))
	}
	if max != math.MaxInt8 {
		t.Errorf("max = %v, want %v", max, int8(math.MaxInt8))
	}
}

func TestMinMaxUint16(t *testing.T) {
	min, max := MinMax[uint16]()
	if min != 0 {
		t.Errorf("min = %v, want 0", min)
	}
	if max != math.MaxUint16 {
		t.Errorf("max = %v, want %v", max, uint16(math.MaxUint16))
	}
}

func TestMinMaxInt16(t *testing.T) {
	min, max := MinMax[int16]()
	if min != math.MinInt16 {
		t.Errorf("min = %v, want %v", min, int16(math.MinInt16))
	}
	if max != math.MaxInt16 {
		t.Errorf("max = %v, want %v", max, int16(math.MaxInt16))
	}
}

func TestMinMaxUint32(t *testing.T) {
	min, max := MinMax[uint32]()
	if min != 0 {
		t.Errorf("min = %v, want 0", min)
	}
	if max != math.MaxUint32 {
		t.Errorf("max = %v, want %v", max, uint32(math.MaxUint32))
	}
}

func TestMinMaxInt32(t *testing.T) {
	min, max := MinMax[int32]()
	if min != math.MinInt32 {
		t.Errorf("min = %v, want %v", min, int32(math.MinInt32))
	}
	if max != math.MaxInt32 {
		t.Errorf("max = %v, want %v", max, int32(math.MaxInt32))
	}
}

func TestMinMaxUint64(t *testing.T) {
	min, max := MinMax[uint64]()
	if min != 0 {
		t.Errorf("min = %v, want 0", min)
	}
	if max != math.MaxUint64 {
		t.Errorf("max = %v, want %v", max, uint64(math.MaxUint64))
	}
}

func TestMinMaxInt64(t *testing.T) {
	min, max := MinMax[int64]()
	if min != math.MinInt64 {
		t.Errorf("min = %v, want %v", min, int64(math.MinInt64))
	}
	if max != math.MaxInt64 {
		t.Errorf("max = %v, want %v", max, int64(math.MaxInt64))
	}
}

func TestMinMaxUintptr(t *testing.T) {
	min, max := MinMax[uintptr]()
	if min != 0 {
		t.Errorf("min = %v, want 0", min)
	}
	if max != ^uintptr(0) {
		t.Errorf("max = %v, want %v", max, ^uintptr(0))
	}
}

func TestMinMaxFloat32(t *testing.T) {
	min, max := MinMax[float32]()
	if min != -math.MaxFloat32 {
		t.Errorf("min = %v, want %v", min, float32(-math.MaxFloat32))
	}
	if max != math.MaxFloat32 {
		t.Errorf("max = %v, want %v", max, float32(math.MaxFloat32))
	}
}

func TestMinMaxFloat64(t *testing.T) {
	min, max := MinMax[float64]()
	if min != -math.MaxFloat64 {
		t.Errorf("min = %v, want %v", min, -math.MaxFloat64)
	}
	if max != math.MaxFloat64 {
		t.Errorf("max = %v, want %v", max, math.MaxFloat64)
	}
}

// byte and rune are aliases for uint8 and int32 respectively, so they share
// those types' case in the switch and should behave identically.
func TestMinMaxByteRuneAliases(t *testing.T) {
	min, max := MinMax[byte]()
	if min != 0 || max != math.MaxUint8 {
		t.Errorf("MinMax[byte]() = (%v, %v), want (0, %v)", min, max, uint8(math.MaxUint8))
	}

	rmin, rmax := MinMax[rune]()
	if rmin != math.MinInt32 || rmax != math.MaxInt32 {
		t.Errorf("MinMax[rune]() = (%v, %v), want (%v, %v)", rmin, rmax, int32(math.MinInt32), int32(math.MaxInt32))
	}
}

func TestMinMaxInt(t *testing.T) {
	min, max := MinMax[int]()
	if min != math.MinInt {
		t.Errorf("min = %v, want %v", min, math.MinInt)
	}
	if max != math.MaxInt {
		t.Errorf("max = %v, want %v", max, math.MaxInt)
	}
}

func TestMinMaxUint(t *testing.T) {
	min, max := MinMax[uint]()
	if min != 0 {
		t.Errorf("min = %v, want 0", min)
	}
	if max != ^uint(0) {
		t.Errorf("max = %v, want %v", max, ^uint(0))
	}
}

// Named types sharing an underlying type with a handled case are not matched
// by the type switch either, since any(zero).(type) checks the dynamic type
// of T exactly, not its underlying type. This documents that MinMax does not
// support user-defined numeric types.
func TestMinMaxNamedTypeUnsupported(t *testing.T) {
	type myInt32 int32

	defer func() {
		if recover() == nil {
			t.Error("MinMax[myInt32]() did not panic; update this test if named types are now supported")
		}
	}()
	MinMax[myInt32]()
}
