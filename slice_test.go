package z_test

import (
	"fmt"
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

		require.Equal(t, dst, []int{4, 9, 16})
	})
	t.Run("same type in-place", func(t *testing.T) {
		src := []int{2, 3, 4}
		z.Map(src, src, func(v int) int {
			return v * v
		})

		require.Equal(t, src, []int{4, 9, 16})
	})
	t.Run("different type", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := []string{"", "", ""}
		z.Map(src, dst, func(v int) string {
			return fmt.Sprintf("%d", v)
		})

		require.Equal(t, dst, []string{"2", "3", "4"})
	})
	t.Run("shorter dst", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := []int{0, 0}
		z.Map(src, dst, func(v int) int {
			return v * v
		})

		require.Equal(t, dst, []int{4, 9})
	})
}

func TestMapped(t *testing.T) {
	t.Run("same type", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := z.Mapped(src, func(v int) int {
			return v * v
		})

		require.Equal(t, dst, []int{4, 9, 16})
	})
	t.Run("different type", func(t *testing.T) {
		src := []int{2, 3, 4}
		dst := z.Mapped(src, func(v int) string {
			return fmt.Sprintf("%d", v)
		})

		require.Equal(t, dst, []string{"2", "3", "4"})
	})
}
