// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestCollect(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3, 4})

	got := seq.Collect()

	assertIntSlice(t, got, []int{1, 2, 3, 4})
}

func TestCollectEmpty(t *testing.T) {
	got := sliceSeq[int](nil).Collect()
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestCollectOrder(t *testing.T) {
	seq := sliceSeq([]string{"z", "a", "m"})

	got := seq.Collect()

	want := []string{"z", "a", "m"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
