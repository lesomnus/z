package z

// Map applies a given operation to each element of the `src` and stores the result in the `dst`.
// The operation is defined by the `op`, which takes an element of type T from the `src` and returns an element of type U to be stored in the `dst`.
// The function processes elements up to the length of the shorter slice between src and dst.
// Returns the `dst` for convenience.
func Map[T any, U any](src []T, dst []U, op func(v T) U) []U {
	MapE(src, dst, func(v T) (U, error) {
		return op(v), nil
	})
	return dst
}

// MapE applies a given operation to each element of the 'src' and stores the result in the 'dst'.
// The operation is defined by the `op`, which takes an element of type T from the `src` and returns an element of type U to be stored in the `dst`.
// The function processes elements up to the length of the shorter slice between src and dst.
// If 'op' returns an error for any element, MapE returns the partially filled 'dst' and the error.
// Otherwise, it returns the fully mapped 'dst' and a nil error.
func MapE[T any, U any](src []T, dst []U, op func(v T) (U, error)) ([]U, error) {
	l := min(len(dst), len(src))

	var err error
	for i := 0; i < l; i++ {
		dst[i], err = op(src[i])
		if err != nil {
			return dst, err
		}
	}

	return dst, nil
}

// Mapped applies a given operation to each element of the `src` and returns a new slice containing the results of the operation.
func Mapped[T any, U any](src []T, op func(v T) U) []U {
	dst := make([]U, len(src))
	return Map(src, dst, op)
}

// Mapped applies a given operation to each element of the `src` and returns a new slice containing the results of the operation.
// If 'op' returns an error for any element, MapE returns the partially filled mapped slice and the error.
func MappedE[T any, U any](src []T, op func(v T) (U, error)) ([]U, error) {
	dst := make([]U, len(src))
	return MapE(src, dst, op)
}
