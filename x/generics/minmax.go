package generics

import (
	"math"

	"golang.org/x/exp/constraints"
)

// MinMax returns the minimum and maximum representable values of T, for
// every integer and float type in the standard library: int8 through
// int64 and uint8 through uint64 (plus int, uint, and uintptr), and float32
// / float64. For float32 and float64, the bounds returned are the most
// negative and most positive finite values (-math.MaxFloat32/64 and
// +math.MaxFloat32/64), not negative/positive infinity.
//
// MinMax determines T's bounds via a type switch on T's exact dynamic
// type, so it only recognizes the types listed above — a named type with
// one of those underlying types (e.g. "type Meters int64") is not matched
// and causes MinMax to panic with "unsupported type", even though such a
// type satisfies the constraints.Integer | constraints.Float constraint.
func MinMax[T constraints.Integer | constraints.Float]() (min T, max T) {
	var zero T
	switch any(zero).(type) {
	case uint8:
		min = any(uint8(0)).(T)
		max = any(uint8(math.MaxUint8)).(T)
	case int8:
		min = any(int8(math.MinInt8)).(T)
		max = any(int8(math.MaxInt8)).(T)
	case uint16:
		min = any(uint16(0)).(T)
		max = any(uint16(math.MaxUint16)).(T)
	case int16:
		min = any(int16(math.MinInt16)).(T)
		max = any(int16(math.MaxInt16)).(T)
	case uint32:
		min = any(uint32(0)).(T)
		max = any(uint32(math.MaxUint32)).(T)
	case int32:
		min = any(int32(math.MinInt32)).(T)
		max = any(int32(math.MaxInt32)).(T)
	case uint64:
		min = any(uint64(0)).(T)
		max = any(uint64(math.MaxUint64)).(T)
	case int64:
		min = any(int64(math.MinInt64)).(T)
		max = any(int64(math.MaxInt64)).(T)
	case uintptr:
		min = any(uintptr(0)).(T)
		max = any(^uintptr(0)).(T)
	case float32:
		min = any(float32(-math.MaxFloat32)).(T)
		max = any(float32(math.MaxFloat32)).(T)
	case float64:
		min = any(float64(-math.MaxFloat64)).(T)
		max = any(float64(math.MaxFloat64)).(T)
	case uint:
		min = any(uint(0)).(T)
		max = any(^uint(0)).(T)
	case int:
		min = any(int(math.MinInt)).(T)
		max = any(int(math.MaxInt)).(T)
	default:
		panic("unsupported type")
	}
	return
}
