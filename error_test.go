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
