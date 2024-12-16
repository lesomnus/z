package z

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
	*a = Fallback[T](*a, b, rest...)
}
