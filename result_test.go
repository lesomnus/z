package z_test

import (
	"io"
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestResult(t *testing.T) {
	t.Run("Take-Into value", func(t *testing.T) {
		v := 0
		err := z.Take(func() (int, error) {
			return 42, nil
		}()).To(&v)
		require.NoError(t, err)
		require.Equal(t, 42, v)
	})
	t.Run("Take-Into error", func(t *testing.T) {
		v := 0
		err := z.Take(func() (int, error) {
			return 36, io.EOF
		}()).To(&v)
		require.ErrorIs(t, err, io.EOF)
		require.Equal(t, 0, v)
	})
	t.Run("IsPending", func(t *testing.T) {
		r := z.Result[int]{}
		require.True(t, r.IsPending())

		r = z.Take(42, nil)
		require.False(t, r.IsPending())

		r = z.Take(0, io.EOF)
		require.False(t, r.IsPending())
	})
	t.Run("Spread", func(t *testing.T) {
		r := z.Take(42, nil)
		v, err := r.Spread()
		require.NoError(t, err)
		require.Equal(t, 42, v)
	})
}
