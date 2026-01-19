package z

type Result[T any] struct {
	Value T
	Error error
}

func Take[T any](v T, err error) Result[T] {
	return Result[T]{Value: v, Error: err}
}

func (r Result[T]) To(p *T) error {
	if r.Error != nil {
		return r.Error
	}
	*p = r.Value
	return nil
}
