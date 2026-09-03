// The unit tests in this file were written by Claude (Anthropic's AI assistant).
package iters

import "testing"

type greeting struct{ name string }

func (g greeting) String() string { return "hello, " + g.name }

func TestToStringOnString(t *testing.T) {
	got := ToString("already a string")
	if got != "already a string" {
		t.Errorf("got %q, want %q", got, "already a string")
	}
}

func TestToStringOnStringer(t *testing.T) {
	got := ToString(greeting{name: "world"})
	if got != "hello, world" {
		t.Errorf("got %q, want %q", got, "hello, world")
	}
}

func TestToStringOnOtherTypes(t *testing.T) {
	if got := ToString(42); got != "42" {
		t.Errorf("ToString(42) = %q, want %q", got, "42")
	}
	if got := ToString(3.14); got != "3.14" {
		t.Errorf("ToString(3.14) = %q, want %q", got, "3.14")
	}
}

func TestStringify(t *testing.T) {
	seq := sliceSeq([]int{1, 2})

	got := collect(Stringify(seq))

	want := []string{"1", "2"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestStringifyMethod(t *testing.T) {
	seq := sliceSeq([]string{"a", "b"})

	got := collect(seq.Stringify())

	want := []string{"a", "b"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestStringifyEmpty(t *testing.T) {
	got := collect(sliceSeq[string](nil).Stringify())
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestStringifyEarlyTermination(t *testing.T) {
	calls := 0
	source := Seq[string](func(yield func(string) bool) {
		for i := 0; i < 100; i++ {
			calls++
			if !yield("x") {
				return
			}
		}
	})

	count := 0
	for range source.Stringify() {
		count++
		if count == 3 {
			break
		}
	}

	if calls != 3 {
		t.Errorf("source yielded %d times, want 3 (early termination not propagated)", calls)
	}
}

func TestJoin(t *testing.T) {
	got := Join(sliceSeq([]string{"a", "b", "c"}), ", ")
	if got != "a, b, c" {
		t.Errorf("got %q, want %q", got, "a, b, c")
	}
}

func TestJoinSingleElement(t *testing.T) {
	got := Join(sliceSeq([]string{"only"}), ", ")
	if got != "only" {
		t.Errorf("got %q, want %q", got, "only")
	}
}

func TestJoinEmpty(t *testing.T) {
	got := Join(sliceSeq[string](nil), ", ")
	if got != "" {
		t.Errorf("got %q, want empty string", got)
	}
}
