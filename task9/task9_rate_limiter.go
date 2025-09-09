package task9

import (
	"context"
	"sync"
	"time"
)

// Task 9: Rate Limiter
// Реализуйте rate limiter с использованием token bucket алгоритма.
// Rate limiter должен ограничивать количество запросов в единицу времени.

// RateLimiter ограничивает количество запросов
type RateLimiter struct {
	tokens int
	burst  int
	mu     sync.Mutex
	ticker *time.Ticker
}

// NewRateLimiter создает новый rate limiter
// rate - количество запросов в секунду
// burst - максимальное количество токенов
func NewRateLimiter(rate int, burst int) *RateLimiter {
	rl := &RateLimiter{
		tokens: burst,
		burst:  burst,
		mu:     sync.Mutex{},
		ticker: nil,
	}
	rl.startRefill(rate)
	return rl
}

func (rl *RateLimiter) startRefill(rate int) {
	var ticker *time.Ticker
	if rate != 0 {
		ticker = time.NewTicker(time.Second * 1 / time.Duration(rate))
	}
	rl.ticker = ticker
	go func() {
		if rl.ticker == nil {
			return
		}
		for range rl.ticker.C {
			rl.mu.Lock()
			rl.tokens++
			if rl.tokens > rl.burst {
				rl.tokens = rl.burst
			}
			rl.mu.Unlock()
		}
	}()
}

// Allow проверяет, можно ли выполнить запрос
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// Wait ждет, пока не станет доступен токен
func (rl *RateLimiter) Wait(ctx context.Context) error {
	ch := make(chan bool)
	go func() {
		for {
			rl.mu.Lock()
			if rl.tokens > 0 {
				ch <- true
				rl.mu.Unlock()
				break
			}
			rl.mu.Unlock()
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ch:
		return nil
	}
}

// Stop останавливает rate limiter
func (rl *RateLimiter) Stop() {
	if rl.ticker != nil {
		rl.ticker.Stop()
	}
}
