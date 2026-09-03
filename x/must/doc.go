// This file's package doc comment was written by Claude (Anthropic's AI assistant).

// Package must converts the (value, error) and (value, value, error) return
// pairs common in Go into a bare value, panicking instead of returning the
// error. It's meant for call sites — initialization code, tests, scripts —
// where an error is truly unexpected and propagating it by hand would only
// add noise.
//
// [Must] takes a (T, error) pair and returns T, panicking if the error is
// non-nil; it composes directly with any function returning that shape, e.g.
// Must(strconv.Atoi("42")). [Must2] does the same for a (T1, T2, error)
// triple, such as os.Pipe's two return values plus its error. [Check] takes
// a bare error and panics if it is non-nil, for call sites that have no
// value to return, only success or failure.
//
// All three panic with an error produced via fmt.Errorf's %w, so the
// original error is still reachable through errors.Is / errors.As / errors.Unwrap
// on the recovered panic value.
package must
