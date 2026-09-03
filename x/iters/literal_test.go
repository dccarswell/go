// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestCreate(t *testing.T) {
	got := collect(Create(1, 2, 3))
	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestCreateEmpty(t *testing.T) {
	got := collect(Create[int]())
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestCreateEarlyTermination(t *testing.T) {
	seq := Create(1, 2, 3, 4, 5)

	var got []int
	seq(func(v int) bool {
		got = append(got, v)
		return len(got) < 2
	})

	assertIntSlice(t, got, []int{1, 2})
}
