// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package generics

import "testing"

func TestSignedTrueForSignedTypes(t *testing.T) {
	if !Signed[int]() {
		t.Error("Signed[int]() = false, want true")
	}
	if !Signed[int8]() {
		t.Error("Signed[int8]() = false, want true")
	}
	if !Signed[int16]() {
		t.Error("Signed[int16]() = false, want true")
	}
	if !Signed[int32]() {
		t.Error("Signed[int32]() = false, want true")
	}
	if !Signed[int64]() {
		t.Error("Signed[int64]() = false, want true")
	}
}

func TestSignedFalseForUnsignedTypes(t *testing.T) {
	if Signed[uint]() {
		t.Error("Signed[uint]() = true, want false")
	}
	if Signed[uint8]() {
		t.Error("Signed[uint8]() = true, want false")
	}
	if Signed[uint16]() {
		t.Error("Signed[uint16]() = true, want false")
	}
	if Signed[uint32]() {
		t.Error("Signed[uint32]() = true, want false")
	}
	if Signed[uint64]() {
		t.Error("Signed[uint64]() = true, want false")
	}
	if Signed[uintptr]() {
		t.Error("Signed[uintptr]() = true, want false")
	}
}

func TestSignedAliases(t *testing.T) {
	if !Signed[rune]() {
		t.Error("Signed[rune]() = false, want true")
	}
	if Signed[byte]() {
		t.Error("Signed[byte]() = true, want false")
	}
}
