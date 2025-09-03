package oncall

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUserRateLimiter(t *testing.T) {
	rl := NewUserLimiter(5, 5)

	const user1 = "user1"
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
}
