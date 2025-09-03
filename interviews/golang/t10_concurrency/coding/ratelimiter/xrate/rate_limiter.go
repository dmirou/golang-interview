package oncall

import (
	"sync"

	"golang.org/x/time/rate"
)

type Limiter interface {
	Allow(key string) bool
}

type UserLimiter struct {
	limiters   map[string]*rate.Limiter
	limitersMu sync.RWMutex

	maxTokens    float64
	refillPerSec float64
}

func NewUserLimiter(maxTokens, refillPerSec float64) *UserLimiter {
	return &UserLimiter{
		limiters:     make(map[string]*rate.Limiter),
		maxTokens:    maxTokens,
		refillPerSec: refillPerSec,
	}
}

func (u *UserLimiter) Allow(key string) bool {
	u.limitersMu.RLock()
	l, ok := u.limiters[key]
	if ok {
		defer u.limitersMu.RUnlock()
		return l.Allow()
	}
	u.limitersMu.RUnlock()

	u.limitersMu.Lock()
	defer u.limitersMu.Unlock()

	u.limiters[key] = rate.NewLimiter(rate.Limit(u.refillPerSec), int(u.maxTokens))

	return u.limiters[key].Allow()
}
