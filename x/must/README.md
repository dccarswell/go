# must

Converts Go's `(value, error)` and `(value, value, error)` return shapes into
a bare value, panicking instead of returning the error. Meant for call sites
where an error is truly unexpected — initialization, tests, scripts — and
threading it through a normal `if err != nil { return err }` would just be
noise.

## Example

```go
package main

import (
	"fmt"
	"os"
	"strconv"

	"go.carswell.tech/x/must"
)

func main() {
	// Must: wraps any (T, error)-returning call directly.
	n := must.Must(strconv.Atoi("42"))
	fmt.Println(n) // 42

	// Must2: wraps any (T1, T2, error)-returning call.
	r, w := must.Must2(os.Pipe())
	defer r.Close()
	defer w.Close()

	// Check: for calls that return only an error.
	must.Check(os.Chdir("."))
}
```

A failing call panics with an error produced via `fmt.Errorf`'s `%w`, so the
original error is still reachable from the recovered panic value with
`errors.Is`, `errors.As`, or `errors.Unwrap`.

## Functions

| Function | Signature                                  | Does |
| -------- | ------------------------------------------- | ---- |
| `Must`   | `func Must[T any](val T, err error) T`      | returns `val`, or panics wrapping `err` |
| `Must2`  | `func Must2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2)` | returns `(v1, v2)`, or panics wrapping `err` |
| `Check`  | `func Check(err error)`                     | panics wrapping `err` if it is non-nil; does nothing otherwise |

---
Portions of this documentation (this README and the package doc comment in
`doc.go`) were written by Claude (Anthropic's AI assistant).
