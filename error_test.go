package z_test

import (
	"errors"
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestWrapErr(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		err := z.WrapErr("a", nil)
		require.NoError(t, err)
	})
	t.Run("non-nil error", func(t *testing.T) {
		err := z.WrapErr("a", errors.New("b"))
		require.Error(t, err)
		require.Equal(t, "a: b", err.Error())
	})
}

func TestCatErr(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		err := z.CatErr("a", nil)
		require.NoError(t, err)
	})
	t.Run("non-nil error", func(t *testing.T) {
		err := z.CatErr("a", errors.New("b"))
		require.Error(t, err)
		require.Equal(t, "ab", err.Error())
	})
}

func TestErr(t *testing.T) {
	t.Run("without args", func(t *testing.T) {
		err := z.Err(errors.New("b"), "a")
		require.Error(t, err)
		require.Equal(t, "a: b", err.Error())
	})
	t.Run("with args", func(t *testing.T) {
		err := z.Err(errors.New("b"), "a %d", 42)
		require.Error(t, err)
		require.Equal(t, "a 42: b", err.Error())
	})
	t.Run("less args", func(t *testing.T) {
		format := "a %d %d"
		err := z.Err(errors.New("b"), format, 42)
		require.Error(t, err)
		require.Equal(t, "a 42 %!d(MISSING): b", err.Error())
	})
	t.Run("more args", func(t *testing.T) {
		format := "a %d"
		err := z.Err(errors.New("b"), format, 42, 36)
		require.Error(t, err)
		require.Equal(t, "a 42%!(EXTRA int=36): b", err.Error())
	})
}

func TestErrIf(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		err := z.ErrIf(nil, "a %d", 42)
		require.NoError(t, err)
	})
	t.Run("non-nil error", func(t *testing.T) {
		err := z.ErrIf(errors.New("b"), "a %d", 42)
		require.Error(t, err)
		require.Equal(t, "a 42: b", err.Error())
	})
}

func TestPanicIf(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		require.NotPanics(t, func() {
			z.PanicIf(nil, "a %d", 42)
		})
	})
	t.Run("non-nil error", func(t *testing.T) {
		require.Panics(t, func() {
			z.PanicIf(errors.New("b"), "a %d", 42)
		})
	})
}
