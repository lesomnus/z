package z_test

import (
	"slices"
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestPredicates(t *testing.T) {
	cases := [][]bool{
		{},

		{true},
		{false},

		{false, false, false},
		{true, false, false},
		{false, true, false},
		{true, true, false},

		{false, false, true},
		{true, false, true},
		{false, true, true},
		{true, true, true},
	}

	t.Run("AllOf", func(t *testing.T) {
		expected := []bool{
			true,

			true,
			false,

			false, false, false, false,
			false, false, false, true,
		}
		for i, c := range cases {
			v := z.AllOf(slices.Values(c), func(v bool) bool { return v })
			require.Equal(t, expected[i], v)
		}
	})
	t.Run("AnyOf", func(t *testing.T) {
		expected := []bool{
			false,

			true,
			false,

			false, true, true, true,
			true, true, true, true,
		}
		for i, c := range cases {
			v := z.AnyOf(slices.Values(c), func(v bool) bool { return v })
			require.Equal(t, expected[i], v)
		}
	})
	t.Run("NoneOf", func(t *testing.T) {
		expected := []bool{
			true,

			false,
			true,

			true, false, false, false,
			false, false, false, false,
		}
		for i, c := range cases {
			v := z.NoneOf(slices.Values(c), func(v bool) bool { return v })
			require.Equal(t, expected[i], v)
		}
	})
}
