package oncall

import (
	"fmt"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type Limiter interface {
	Allow(key string) bool
}

type RateLimiter struct {
	tokens      float64
	tokensMutex sync.RWMutex

	maxTokens    float64
	refillPerSec float64
	lastRefill   time.Time

	quit chan struct{}
}

func NewRateLimiter(maxTokens float64, refillPerSec float64) *RateLimiter {
	rl := RateLimiter{
		maxTokens:    maxTokens,
		refillPerSec: refillPerSec,
		quit:         make(chan struct{}),
	}
	rl.refillTokens()

	return &rl
}

func (r *RateLimiter) refillTokens() {
	defer func() {
		fmt.Println("refill tokens done", "tokens", r.tokens)
	}()

	now := time.Now()

	if r.lastRefill.IsZero() {
		r.tokens = min(r.refillPerSec, r.maxTokens)
	} else {
		diff := decimal.NewFromFloat(r.refillPerSec).
			Mul(decimal.NewFromFloat(now.Sub(r.lastRefill).Seconds()))

		newTokens, _ := decimal.NewFromFloat(r.tokens).Add(diff).Float64()
		r.tokens = min(newTokens, r.maxTokens)
	}
	r.lastRefill = now
}

func (r *RateLimiter) Allow() bool {
	r.tokensMutex.Lock()
	defer r.tokensMutex.Unlock()

	r.refillTokens()

	if r.tokens >= 1 {
		r.tokens -= 1
		return true
	}

	return false
}

type UserLimiter struct {
	limiters   map[string]*RateLimiter
	limitersMu sync.RWMutex

	maxTokens    float64
	refillPerSec float64
}

func NewUserLimiter(maxTokens, refillPerSec float64) *UserLimiter {
	return &UserLimiter{
		limiters:     make(map[string]*RateLimiter),
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

	u.limiters[key] = NewRateLimiter(u.maxTokens, u.refillPerSec)
	return u.limiters[key].Allow()
}
