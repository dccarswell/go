package generics

// Ptr returns a pointer to a new variable holding v. It is useful for
// taking the address of a value that isn't itself addressable, such as a
// literal or the result of an expression (e.g. Ptr(5) or Ptr(x+1), neither
// of which can be written as &5 or &(x+1)).
func Ptr[T any](v T) *T {
	return &v
}
