// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func sliceSeq[T any](items []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range items {
			if !yield(v) {
				return
			}
		}
	}
}

func sliceSeq2[K comparable, V any](keys []K, vals []V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := range keys {
			if !yield(keys[i], vals[i]) {
				return
			}
		}
	}
}

func TestSeqRangeOverFunc(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	var got []int
	for v := range seq {
		got = append(got, v)
	}

	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestSeq2RangeOverFunc(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	var keys []string
	var vals []int
	for k, v := range seq {
		keys = append(keys, k)
		vals = append(vals, v)
	}

	wantKeys := []string{"a", "b"}
	wantVals := []int{1, 2}
	if len(keys) != len(wantKeys) {
		t.Fatalf("keys = %v, want %v", keys, wantKeys)
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] || vals[i] != wantVals[i] {
			t.Errorf("pair[%d] = (%s, %d), want (%s, %d)", i, keys[i], vals[i], wantKeys[i], wantVals[i])
		}
	}
}

func TestPull(t *testing.T) {
	seq := sliceSeq([]int{10, 20, 30})
	next, stop := Pull(seq)
	defer stop()

	for _, want := range []int{10, 20, 30} {
		got, ok := next()
		if !ok {
			t.Fatalf("next() = (_, false), want (%d, true)", want)
		}
		if got != want {
			t.Errorf("next() = %d, want %d", got, want)
		}
	}

	if _, ok := next(); ok {
		t.Error("next() after exhaustion returned ok=true, want false")
	}
}

func TestPullEmpty(t *testing.T) {
	next, stop := Pull(sliceSeq[int](nil))
	defer stop()

	if _, ok := next(); ok {
		t.Error("next() on empty sequence returned ok=true, want false")
	}
}

func TestPullStop(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4, 5})
	next, stop := Pull(seq)

	got, ok := next()
	if !ok || got != 1 {
		t.Fatalf("next() = (%d, %v), want (1, true)", got, ok)
	}

	stop()

	if _, ok := next(); ok {
		t.Error("next() after stop() returned ok=true, want false")
	}
}

func TestPull2(t *testing.T) {
	seq := sliceSeq2([]string{"x", "y"}, []int{100, 200})
	next, stop := Pull2(seq)
	defer stop()

	k, v, ok := next()
	if !ok || k != "x" || v != 100 {
		t.Fatalf("next() = (%s, %d, %v), want (x, 100, true)", k, v, ok)
	}

	k, v, ok = next()
	if !ok || k != "y" || v != 200 {
		t.Fatalf("next() = (%s, %d, %v), want (y, 200, true)", k, v, ok)
	}

	if _, _, ok := next(); ok {
		t.Error("next() after exhaustion returned ok=true, want false")
	}
}

func TestPull2Empty(t *testing.T) {
	next, stop := Pull2(sliceSeq2[string, int](nil, nil))
	defer stop()

	if _, _, ok := next(); ok {
		t.Error("next() on empty sequence returned ok=true, want false")
	}
}
