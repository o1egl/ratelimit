package ratelimit

import (
	"context"
	"sync"
	"time"
)

// RateLimiter
type RateLimiter interface {
	// Allow returns true if limit is not reached
	Allow() bool
	// WaitFor returns time to appearance of new resources
	WaitFor() time.Duration
	// Get gets available resource or waits for appearance of new resources
	Get(ctx context.Context) error
}

// NewRatelimiter returns new instance of RateLimiter
func NewRatelimiter(limit int, interval time.Duration) RateLimiter {
	return &rateLimit{limit: limit, interval: interval}
}

type rateLimit struct {
	sync.RWMutex
	limit            int
	interval         time.Duration
	requestsCount    int
	firstRequestTime time.Time
}

// Allow implements RateLimiter.Allow
func (r *rateLimit) Allow() bool {
	r.Lock()
	defer r.Unlock()
	if time.Now().After(r.firstRequestTime) {
		r.requestsCount = 0
	}
	switch r.requestsCount {
	case 0:
		r.requestsCount++
		r.firstRequestTime = time.Now().Add(r.interval)
		return true
	case r.limit:
		return false
	default:
		r.requestsCount++
		return true
	}
}

// Allow implements RateLimiter.WaitFor
func (r *rateLimit) WaitFor() time.Duration {
	r.RLock()
	defer r.RUnlock()
	return r.firstRequestTime.Sub(time.Now())
}

func (r *rateLimit) Get(ctx context.Context) error {
	for !r.Allow() {
		select {
		case <-time.After(r.WaitFor()):
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
