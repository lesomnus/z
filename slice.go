package z

// Map applies a given operation to each element of the `src` and stores the result in the `dst` slice.
// The operation is defined by the `op`, which takes an element of type T from the `src` and returns an element of type U to be stored in the `dst`.
// The function processes elements up to the length of the shorter slice between src and dst.
// Returns the `dst` for con
func Map[T any, U any](src []T, dst []U, op func(v T) U) []U {
	l := len(src)
	if len(dst) < l {
		l = len(dst)
	}

	for i := 0; i < l; i++ {
		dst[i] = op(src[i])
	}

	return dst
}

// Mapped applies a given operation to each element of the `src` slice and returns a new slice containing the results of the operation.
func Mapped[T any, U any](src []T, op func(v T) U) []U {
	dst := make([]U, len(src))
	Map(src, dst, op)

	return dst
}
