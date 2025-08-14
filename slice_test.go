package z_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Run("same type", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := []int{0, 0, 0}
		z.Map(src, dst, func(v int) int {
			return v * v
		})

		require.Equal(t, []int{4, 9, 16}, dst)
	})
	t.Run("same type in-place", func(t *testing.T) {
		src := []int{2, 3, 4}
		z.Map(src, src, func(v int) int {
			return v * v
		})

		require.Equal(t, []int{4, 9, 16}, src)
	})
	t.Run("different type", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := []string{"", "", ""}
		z.Map(src, dst, func(v int) string {
			return fmt.Sprintf("%d", v)
		})

		require.Equal(t, []string{"2", "3", "4"}, dst)
	})
	t.Run("shorter dst", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := []int{0, 0}
		z.Map(src, dst, func(v int) int {
			return v * v
		})

		require.Equal(t, []int{4, 9}, dst)
	})
}

func TestMapE(t *testing.T) {
	t.Run("error is forwarded", func(t *testing.T) {
		err := errors.New("foo")
		vs := []int{2, 3, 4}
		_, err_received := z.MapE(vs, vs, func(v int) (int, error) {
			return 0, err
		})
		require.Same(t, err, err_received)
	})
	t.Run("partially filled", func(t *testing.T) {
		src := []int{2, 3, 4, 5, 6}
		dst := []int{0, 0, 0, 0, 0}
		z.MapE(src, dst, func(v int) (int, error) {
			if v > 3 {
				return 42, io.EOF
			}

			return v, nil
		})

		require.Equal(t, []int{2, 3, 42, 0, 0}, dst)
	})
}

func TestMapped(t *testing.T) {
	t.Run("same type", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := z.Mapped(src, func(v int) int {
			return v * v
		})

		require.Equal(t, []int{4, 9, 16}, dst)
	})
	t.Run("different type", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := z.Mapped(src, func(v int) string {
			return fmt.Sprintf("%d", v)
		})

		require.Equal(t, []string{"2", "3", "4"}, dst)
	})
}
