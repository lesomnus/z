package z

import "sync"

type Exclusive[T any] struct {
	m sync.Mutex
	v T
}

func NewExclusive[T any](v T) *Exclusive[T] {
	return &Exclusive[T]{v: v}
}

func (e *Exclusive[T]) Lock() (ExclusiveValue[T], func()) {
	e.m.Lock()
	return ExclusiveValue[T]{&e.v}, e.m.Unlock
}

func (e *Exclusive[T]) Use(f func(v ExclusiveValue[T])) {
	v, unlock := e.Lock()
	defer unlock()

	f(v)
}

type ExclusiveValue[T any] struct {
	v *T
}

func (c ExclusiveValue[T]) Get() T {
	return *c.v
}

func (c ExclusiveValue[T]) Set(v T) {
	*c.v = v
}
