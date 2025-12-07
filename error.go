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

// Err returns a formatted error wrapping `err`.
//
//	fmt.Errorf(format+": %w", args..., err)
func Err(err error, format string, args ...any) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %w", msg, err)
}

// ErrIf returns nil if `err` is nil.
// Otherwise it returns:
//
//	fmt.Errorf(format+": %w", args..., err)
func ErrIf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return Err(err, format, args...)
}

// PanicIf panics with a formatted message if `err` is not nil.
func PanicIf(err error, msg string, args ...any) {
	if err != nil {
		panic(fmt.Errorf(msg+": %w", append(args, err)...))
	}
}
