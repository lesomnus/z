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

// Filter appends elements of `src` that satisfy `p` into `dst` and returns the resulting `dst`.
// For correctness, `dst` should not overlap `src`'s underlying array.
func Filter[T any](src []T, dst []T, p func(v T) bool) []T {
	for i := range src {
		if !p(src[i]) {
			continue
		}

		dst = append(dst, src[i])
	}

	return dst
}

// Filtered returns a newly allocated slice containing elements of `src` that satisfy `p`.
// The relative order of kept elements is preserved.
func Filtered[T any](src []T, p func(v T) bool) []T {
	dst := make([]T, 0, len(src))
	return Filter(src, dst, p)
}

// FilterInPlace partitions `src` in-place and returns the prefix containing elements that satisfy `p`.
// Elements rejected by `p` are moved into the tail of the original slice (i.e., `src[len(result):]`).
// The relative order of elements is not preserved.
func FilterInPlace[T any](src []T, p func(v T) bool) []T {
	j := len(src)
	for i := 0; i < j; {
		if p(src[i]) {
			i++
			continue
		}

		j--
		src[i], src[j] = src[j], src[i]
	}

	return src[:j]
}

// FilteredInPlace partitions `src` in-place and returns two slices: kept and rejected.
// Both returned slices share the same underlying array as `src`.
func FilteredInPlace[T any](src []T, p func(v T) bool) ([]T, []T) {
	rst := FilterInPlace(src, p)
	return rst, src[len(rst):]
}
