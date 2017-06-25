package ratelimit

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimit_Get(t *testing.T) {
	const (
		limit    = 100
		interval = time.Second
	)

	var (
		counter        int64
		ctx, ctxCancel = context.WithCancel(context.Background())
		start          = make(chan bool)
		wg             = sync.WaitGroup{}
		r              = NewRatelimiter(limit, interval)
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for {
				if err := r.Get(ctx); err == nil {
					atomic.AddInt64(&counter, 1)
					continue
				}
				select {
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	close(start)
	<-time.After(3 * time.Second)
	ctxCancel()
	wg.Wait()

	if 300 != counter {
		t.Errorf("expected: 300, obtained: %d", counter)
	}
}
