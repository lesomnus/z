package z

import (
	"context"
	"fmt"
	"reflect"
	"sync/atomic"
)

type ctxKey uint64

var ctxKeyV atomic.Uint64

type Use[T any] struct {
	k ctxKey
}

func NewUse[T any]() Use[T] {
	k := ctxKeyV.Add(1)
	return Use[T]{k: ctxKey(k)}
}

func (u Use[T]) Into(ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, u.k, v)
}

func (u Use[T]) From(ctx context.Context) (T, bool) {
	v, ok := ctx.Value(u.k).(T)
	return v, ok
}

func (u Use[T]) Must(ctx context.Context) T {
	v, ok := u.From(ctx)
	if !ok {
		panic(fmt.Sprintf("no %v is provided", reflect.TypeOf(v).Name()))
	}

	return v
}
