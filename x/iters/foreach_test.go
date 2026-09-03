// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestForEach(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	var got []int
	ForEach(func(v int) { got = append(got, v) })(seq)

	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestForEachMethod(t *testing.T) {
	seq := sliceSeq([]int{1, 2, 3})

	var got []int
	seq.ForEach(func(v int) { got = append(got, v) })

	assertIntSlice(t, got, []int{1, 2, 3})
}

func TestForEachEmpty(t *testing.T) {
	calls := 0
	sliceSeq[int](nil).ForEach(func(int) { calls++ })
	if calls != 0 {
		t.Errorf("callback invoked %d times on empty sequence, want 0", calls)
	}
}

func TestForEach2(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	var keys []string
	var vals []int
	ForEach2(func(k string, v int) {
		keys = append(keys, k)
		vals = append(vals, v)
	})(seq)

	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" || vals[0] != 1 || vals[1] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[a b] vals=[1 2]", keys, vals)
	}
}

func TestForEach2Method(t *testing.T) {
	seq := sliceSeq2([]string{"a", "b"}, []int{1, 2})

	var keys []string
	var vals []int
	seq.ForEach(func(k string, v int) {
		keys = append(keys, k)
		vals = append(vals, v)
	})

	if len(keys) != 2 || keys[0] != "a" || keys[1] != "b" || vals[0] != 1 || vals[1] != 2 {
		t.Errorf("keys=%v vals=%v, want keys=[a b] vals=[1 2]", keys, vals)
	}
}

func TestForEach2Empty(t *testing.T) {
	calls := 0
	sliceSeq2[string, int](nil, nil).ForEach(func(string, int) { calls++ })
	if calls != 0 {
		t.Errorf("callback invoked %d times on empty sequence, want 0", calls)
	}
}
