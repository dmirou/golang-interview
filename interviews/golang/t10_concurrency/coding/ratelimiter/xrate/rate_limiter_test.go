package oncall

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUserRateLimiter(t *testing.T) {
	const user1 = "user1"

	ticker := time.NewTicker(time.Second)
	rl := NewUserLimiter(10, 5)

	for i := 0; i < 5; i++ {
		select {
		case <-ticker.C:
			t.Log("test ticker tick")
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))
			require.True(t, rl.Allow(user1))
		}
	}
	ticker.Stop()
}
