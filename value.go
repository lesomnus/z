package z

// Ptr returns pointer to the temporal variable which value is equal to `v`.
// This function is intended to use with the constants or literals.
//
//	v := Ptr(42) // Do
//	v := Ptr(u) // Don't
//	v := &u // Do
//
// Be careful with that returned pointer does not pointing `v` itself.
//
//	&v != Ptr(v)
func Ptr[T any](v T) *T {
	return &v
}

// Fallback selects the first non-zero value from the given values.
//
//	Fallback(0, 0, 42)  // 42
//	Fallback("", "foo") // "foo"
func Fallback[T comparable](a, b T, rest ...T) T {
	var z T
	if a != z {
		return a
	}
	if b != z {
		return b
	}
	for _, c := range rest {
		if c != z {
			return c
		}
	}
	return z
}

// FallbackP selects the first non-zero value from the given values and set to `*a`.
// It panics if `a` is nil. The value of `*a` not updated if `*a` is non-zero value.
func FallbackP[T comparable](a *T, b T, rest ...T) {
	if a == nil {
		panic("target must not nil")
	}
	*a = Fallback(*a, b, rest...)
}
