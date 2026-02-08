package z

type Result[T any] struct {
	Value *T
	Error error
}

func Take[T any](v T, err error) Result[T] {
	return Result[T]{Value: &v, Error: err}
}

func (r Result[T]) To(p *T) error {
	if r.Error != nil {
		return r.Error
	}
	if r.Value != nil {
		*p = *r.Value
	}
	return nil
}

func (r Result[T]) IsPending() bool {
	return r.Value == nil && r.Error == nil
}

func (r Result[T]) Spread() (T, error) {
	var z T
	if r.Error != nil {
		return z, r.Error
	}
	if r.Value == nil {
		return z, nil
	}
	return *r.Value, nil
}

func (r Result[T]) Must() T {
	if r.Error != nil {
		panic(r.Error)
	}
	if r.Value == nil {
		panic("result is pending")
	}

	return *r.Value
}
