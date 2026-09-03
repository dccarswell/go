// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import (
	"errors"
	"testing"
)

func TestWrapStructOk(t *testing.T) {
	if !(WrapStruct[int]{Value: 1}).Ok() {
		t.Error("WrapStruct with nil Error: Ok() = false, want true")
	}
	if (WrapStruct[int]{Error: errors.New("boom")}).Ok() {
		t.Error("WrapStruct with non-nil Error: Ok() = true, want false")
	}
}

func TestWrapStruct2Ok(t *testing.T) {
	if !(WrapStruct2[string, int]{Key: "k"}).Ok() {
		t.Error("WrapStruct2 with nil Error: Ok() = false, want true")
	}
	ws := WrapStruct2[string, int]{Key: "k"}
	ws.Error = errors.New("boom")
	if ws.Ok() {
		t.Error("WrapStruct2 with non-nil Error: Ok() = true, want false")
	}
}

func halveOrErr(v int) WrapStruct[int] {
	if v%2 != 0 {
		return WrapStruct[int]{Error: errors.New("odd")}
	}
	return WrapStruct[int]{Value: v / 2}
}

func TestWrap(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4})

	got := collect(Wrap(WrapFunc[int](halveOrErr))(seq))

	if len(got) != 4 {
		t.Fatalf("got %d results, want 4", len(got))
	}
	wantOk := []bool{false, true, false, true}
	wantVal := []int{0, 1, 0, 2}
	for i := range got {
		if got[i].Ok() != wantOk[i] {
			t.Errorf("got[%d].Ok() = %v, want %v", i, got[i].Ok(), wantOk[i])
		}
		if got[i].Ok() && got[i].Value != wantVal[i] {
			t.Errorf("got[%d].Value = %d, want %d", i, got[i].Value, wantVal[i])
		}
	}
}

func TestWrapMethod(t *testing.T) {
	seq := sliceSeq([]int{2, 4})

	got := collect(seq.Wrap(WrapFunc[int](halveOrErr)))

	if len(got) != 2 {
		t.Fatalf("got %d results, want 2", len(got))
	}
	for i, want := range []int{1, 2} {
		if !got[i].Ok() {
			t.Fatalf("got[%d].Ok() = false, want true", i)
		}
		if got[i].Value != want {
			t.Errorf("got[%d].Value = %d, want %d", i, got[i].Value, want)
		}
	}
}

func TestWrapEmpty(t *testing.T) {
	got := collect(sliceSeq[int](nil).Wrap(WrapFunc[int](halveOrErr)))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestWrapEarlyTermination(t *testing.T) {
	calls := 0
	source := Seq[int](func(yield func(int) bool) {
		for i := 0; i < 100; i += 2 {
			calls++
			if !yield(i) {
				return
			}
		}
	})

	count := 0
	for range source.Wrap(WrapFunc[int](halveOrErr)) {
		count++
		if count == 3 {
			break
		}
	}

	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (early termination not propagated)", calls)
	}
}

func pairOrErr(k string, v int) WrapStruct2[string, int] {
	if v < 0 {
		return WrapStruct2[string, int]{Key: k, WrapStruct: WrapStruct[int]{Error: errors.New("negative")}}
	}
	return WrapStruct2[string, int]{Key: k, WrapStruct: WrapStruct[int]{Value: v}}
}

func TestWrap2(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c"}, []int{1, -1, 3})

	got := collect(Wrap2(WrapFunc2[string, int](pairOrErr))(seq))

	if len(got) != 3 {
		t.Fatalf("got %d results, want 3", len(got))
	}
	wantKey := []string{"a", "b", "c"}
	wantOk := []bool{true, false, true}
	wantVal := []int{1, 0, 3}
	for i := range got {
		if got[i].Key != wantKey[i] {
			t.Errorf("got[%d].Key = %q, want %q", i, got[i].Key, wantKey[i])
		}
		if got[i].Ok() != wantOk[i] {
			t.Errorf("got[%d].Ok() = %v, want %v", i, got[i].Ok(), wantOk[i])
		}
		if got[i].Ok() && got[i].Value != wantVal[i] {
			t.Errorf("got[%d].Value = %d, want %d", i, got[i].Value, wantVal[i])
		}
	}
}

func TestWrap2Method(t *testing.T) {
	seq := sliceSeq2([]string{"x", "y"}, []int{10, 20})

	got := collect(seq.Wrap2(WrapFunc2[string, int](pairOrErr)))

	if len(got) != 2 {
		t.Fatalf("got %d results, want 2", len(got))
	}
	for i, want := range []struct {
		key string
		val int
	}{{"x", 10}, {"y", 20}} {
		if !got[i].Ok() {
			t.Fatalf("got[%d].Ok() = false, want true", i)
		}
		if got[i].Key != want.key || got[i].Value != want.val {
			t.Errorf("got[%d] = (%q, %d), want (%q, %d)", i, got[i].Key, got[i].Value, want.key, want.val)
		}
	}
}

func TestWrap2Empty(t *testing.T) {
	got := collect(sliceSeq2[string, int](nil, nil).Wrap2(WrapFunc2[string, int](pairOrErr)))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestWrap2EarlyTermination(t *testing.T) {
	calls := 0
	source := Seq2[string, int](func(yield func(string, int) bool) {
		for i := 0; i < 100; i++ {
			calls++
			if !yield("k", i) {
				return
			}
		}
	})

	count := 0
	for range source.Wrap2(WrapFunc2[string, int](pairOrErr)) {
		count++
		if count == 3 {
			break
		}
	}

	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (early termination not propagated)", calls)
	}
}
