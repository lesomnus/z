package z_test

import (
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestExpectOneOf(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		err := z.ExpectOneOf(42, 1, 2, 42, 3)
		require.NoError(t, err)
	})
	t.Run("not contains", func(t *testing.T) {
		err := z.ExpectOneOf(42, 1, 2, 3)
		require.Error(t, err)
		require.ErrorContains(t, err, "1, 2, 3")
	})
}
