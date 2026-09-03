// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestCoalesce(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b", "c"}, []int{1, 2, 3})
	join := func(k string, v int) string {
		switch v {
		case 1:
			return "a1"
		case 2:
			return "b2"
		default:
			return "c3"
		}
	}

	got := collect(Coalesce(CoalesceFunc[string, int, string](join))(seq))

	want := []string{"a1", "b2", "c3"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestCoalesceMethod(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	got := collect(seq.Coalesce(func(k string, v int) int { return v * 10 }))
	assertIntSlice(t, got, []int{10, 20})
}

func TestCoalesceEmpty(t *testing.T) {
	got := collect(sliceSeq2[string, int](nil, nil).Coalesce(func(k string, v int) int { return v }))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestCoalesceEarlyTermination(t *testing.T) {
	calls := 0
	source := Seq2[int, int](func(yield func(int, int) bool) {
		for i := 1; i <= 100; i++ {
			calls++
			if !yield(i, i) {
				return
			}
		}
	})

	count := 0
	for range source.Coalesce(func(k, v int) int { return k + v }) {
		count++
		if count == 3 {
			break
		}
	}

	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (early termination not propagated)", calls)
	}
}
