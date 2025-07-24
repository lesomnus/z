package z_test

import (
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestUse(t *testing.T) {
	use_a := z.NewUse[int]()
	use_b := z.NewUse[string]()
	use_c := z.NewUse[string]()

	ctx_a := use_a.Into(t.Context(), 42)
	ctx_b := use_b.Into(ctx_a, "b")
	ctx_c := use_c.Into(ctx_b, "c")

	require.Equal(t, 42, use_a.Must(ctx_a))
	require.Equal(t, 42, use_a.Must(ctx_b))
	require.Equal(t, 42, use_a.Must(ctx_c))
	require.Equal(t, "b", use_b.Must(ctx_b))
	require.Equal(t, "b", use_b.Must(ctx_c))
	require.Equal(t, "c", use_c.Must(ctx_c))

	v := ctx_a.Value(uint64(0))
	require.IsType(t, nil, v)
	v = ctx_b.Value(uint64(0))
	require.IsType(t, nil, v)
	v = ctx_c.Value(uint64(0))
	require.IsType(t, nil, v)
}
