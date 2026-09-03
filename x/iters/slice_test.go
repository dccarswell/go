// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

func TestValues(t *testing.T) {
	got := collect(Values([]int{1, 2, 3, 4}))
	assertIntSlice(t, got, []int{1, 2, 3, 4})
}

func TestValuesOrder(t *testing.T) {
	got := collect(Values([]string{"z", "a", "m"}))

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

func TestValuesEmpty(t *testing.T) {
	got := collect(Values[int](nil))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestValuesEarlyTermination(t *testing.T) {
	seq := Values([]int{1, 2, 3, 4, 5})

	var got []int
	seq(func(v int) bool {
		got = append(got, v)
		return len(got) < 2
	})

	assertIntSlice(t, got, []int{1, 2})
}

func TestValuesRoundTripsWithCollect(t *testing.T) {
	want := []int{5, 4, 3, 2, 1}

	got := Values(want).Collect()

	assertIntSlice(t, got, want)
}

func TestUnslice(t *testing.T) {
	keys, vals := collect2(Unslice([]string{"a", "b", "c"}))

	wantKeys := []int{0, 1, 2}
	wantVals := []string{"a", "b", "c"}
	if len(keys) != len(wantKeys) {
		t.Fatalf("keys=%v, want %v", keys, wantKeys)
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] || vals[i] != wantVals[i] {
			t.Errorf("pair[%d] = (%d, %q), want (%d, %q)", i, keys[i], vals[i], wantKeys[i], wantVals[i])
		}
	}
}

func TestUnsliceEmpty(t *testing.T) {
	keys, vals := collect2(Unslice[string](nil))
	if len(keys) != 0 || len(vals) != 0 {
		t.Errorf("keys=%v vals=%v, want empty", keys, vals)
	}
}

func TestUnsliceEarlyTermination(t *testing.T) {
	seq := Unslice([]string{"a", "b", "c", "d", "e"})

	var gotKeys []int
	seq(func(k int, v string) bool {
		gotKeys = append(gotKeys, k)
		return len(gotKeys) < 2
	})

	assertIntSlice(t, gotKeys, []int{0, 1})
}

func TestUnmap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	got := make(map[string]int)
	for k, v := range Unmap(m) {
		got[k] = v
	}

	if len(got) != len(m) {
		t.Fatalf("got %v, want %v", got, m)
	}
	for k, want := range m {
		if got[k] != want {
			t.Errorf("got[%q] = %d, want %d", k, got[k], want)
		}
	}
}

func TestUnmapEmpty(t *testing.T) {
	got := make(map[string]int)
	for k, v := range Unmap[string, int](nil) {
		got[k] = v
	}
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestUnmapEarlyTermination(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	seq := Unmap(m)

	count := 0
	seq(func(k string, v int) bool {
		count++
		return count < 2
	})

	if count != 2 {
		t.Errorf("yielded %d times before stopping, want 2", count)
	}
}
