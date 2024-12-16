package z

import "fmt"

// WrapErr returns `err` as-is if `err` is nil.
// Otherwise it returns:
//
//	fmt.Errorf("%s: %w", msg, err)
func WrapErr(msg string, err error) error {
	if err == nil {
		return err
	}

	return fmt.Errorf("%s: %w", msg, err)
}

// CatErr returns `err` as-is if `err` is nil.
// Otherwise it returns:
//
//	fmt.Errorf("%s%w", msg, err)
func CatErr(msg string, err error) error {
	if err == nil {
		return err
	}

	return fmt.Errorf("%s%w", msg, err)
}
