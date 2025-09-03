package ontick

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(5, 5)
	done := make(chan struct{})

	go func() {
		rl.Run()
		done <- struct{}{}
	}()

	const user1 = "user1"
	time.Sleep(100 * time.Millisecond)
	ticker := time.NewTicker(time.Second)

	for i := 0; i < 5; i++ {
		select {
		case <-ticker.C:
			t.Log("test ticker tick")
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))

			require.False(t, rl.Allow(user1))
			require.False(t, rl.Allow(user1))
		}
	}
	ticker.Stop()
	rl.Stop()
	<-done
}
