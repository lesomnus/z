package z_test

import (
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	t.Run("literal string", func(t *testing.T) {
		v := z.Ptr("foo")
		require.Equal(t, "foo", *v)
	})
	t.Run("string variable", func(t *testing.T) {
		u := "foo"
		v := z.Ptr(u)
		require.Equal(t, u, *v)
		require.NotSame(t, &u, v)
	})
}

func TestFallback(t *testing.T) {
	t.Run("zero2", func(t *testing.T) {
		v := z.Fallback(0, 0)
		require.Equal(t, 0, v)
	})
	t.Run("zero3", func(t *testing.T) {
		v := z.Fallback(0, 0, 0)
		require.Equal(t, 0, v)
	})
	t.Run("zero4", func(t *testing.T) {
		v := z.Fallback(0, 0, 0, 0)
		require.Equal(t, 0, v)
	})
	t.Run("int2", func(t *testing.T) {
		v := z.Fallback(0, 42)
		require.Equal(t, 42, v)
	})
	t.Run("int3", func(t *testing.T) {
		v := z.Fallback(0, 0, 42)
		require.Equal(t, 42, v)
	})
	t.Run("int4", func(t *testing.T) {
		v := z.Fallback(0, 0, 0, 42)
		require.Equal(t, 42, v)
	})
	t.Run("first", func(t *testing.T) {
		v := z.Fallback(42, 0)
		require.Equal(t, 42, v)
	})
	t.Run("middle", func(t *testing.T) {
		v := z.Fallback(0, 0, 42, -1)
		require.Equal(t, 42, v)
	})
}

func TestFallbackP(t *testing.T) {
	t.Run("non-zero", func(t *testing.T) {
		v := 36
		z.FallbackP(&v, 42)
		require.Equal(t, 36, v)
	})
	t.Run("zero", func(t *testing.T) {
		v := 0
		z.FallbackP(&v, 42)
		require.Equal(t, 42, v)
	})
	t.Run("middle", func(t *testing.T) {
		v := 0
		z.FallbackP(&v, 0, 42, -1)
		require.Equal(t, 42, v)
	})
	t.Run("nil target", func(t *testing.T) {
		require.Panics(t, func() {
			z.FallbackP(nil, 0)
		})
	})
}
