package z

import (
	"fmt"
	"slices"
	"strings"
)

// ExpectOneOf returns nil if `vs` contains `v`.
// Otherwise, it returns an error with a message indicating that `vs` does not contain `v`.
func ExpectOneOf[T comparable](v T, vs ...T) error {
	if slices.Contains(vs, v) {
		return nil
	}

	o := make([]string, len(vs))
	for i, v := range vs {
		o[i] = fmt.Sprintf("%v", v)
	}

	return fmt.Errorf(`expected the value to be one of [%s], but it was [%v]`, strings.Join(o, ", "), v)
}
