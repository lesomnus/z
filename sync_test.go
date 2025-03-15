package z_test

import (
	"sync"
	"testing"

	"github.com/lesomnus/z"
	"github.com/stretchr/testify/require"
)

func TestExclusive(t *testing.T) {
	t.Run("access", func(t *testing.T) {
		e := z.NewExclusive(42)

		ev, unlock := e.Lock()
		defer unlock()

		v := ev.Get()
		require.Equal(t, 42, v)

		ev.Set(36)
		v = ev.Get()
		require.Equal(t, 36, v)
	})
	t.Run("race", func(t *testing.T) {
		e := z.NewExclusive(0)

		const N = 1000
		var wg sync.WaitGroup
		wg.Add(N)
		for range N {
			go func() {
				defer wg.Done()
				e.Use(func(v z.ExclusiveValue[int]) {
					v.Set(v.Get() + 1)
				})
			}()
		}

		wg.Wait()

		ev, unlock := e.Lock()
		defer unlock()

		v := ev.Get()
		require.Equal(t, N, v)
	})
}
