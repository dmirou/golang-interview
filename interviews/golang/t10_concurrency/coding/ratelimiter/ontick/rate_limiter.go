package ontick

import (
	"fmt"
	"sync"
	"time"
)

type Limiter interface {
	Allow(key string) bool
}

type RateLimiter struct {
	tokens      map[string]int
	tokensMutex sync.RWMutex

	maxTokens    int
	refillPerSec int

	quit chan struct{}
}

func NewRateLimiter(maxTokens int, refillPerSec int) *RateLimiter {
	return &RateLimiter{
		tokens:       make(map[string]int),
		maxTokens:    maxTokens,
		refillPerSec: refillPerSec,
		quit:         make(chan struct{}),
	}
}

func (r *RateLimiter) Run() {
	ticker := time.NewTicker(time.Second)
	defer func() {
		ticker.Stop()
		r.quit <- struct{}{}
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("rate limiter ticket tick")
			r.tokensMutex.Lock()
			for k, v := range r.tokens {
				r.tokens[k] = min(v+r.refillPerSec, r.maxTokens)
			}
			r.tokensMutex.Unlock()
		case <-r.quit:
			return
		}
	}
}

func (r *RateLimiter) Stop() {
	defer fmt.Println("rate limiter stopped")
	r.quit <- struct{}{}
	<-r.quit
}

func (r *RateLimiter) Allow(key string) bool {
	r.tokensMutex.Lock()
	defer r.tokensMutex.Unlock()

	tokens, ok := r.tokens[key]
	if !ok {
		r.tokens[key] = r.refillPerSec
		tokens = r.refillPerSec
	}
	if tokens <= 0 {
		return false
	}

	r.tokens[key] = tokens - 1

	return true
}
