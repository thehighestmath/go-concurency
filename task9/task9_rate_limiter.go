package task9

import (
	"context"
	"fmt"
	"time"
)

// Task 9: Rate Limiter
// Реализуйте rate limiter с использованием token bucket алгоритма.
// Rate limiter должен ограничивать количество запросов в единицу времени.

// RateLimiter ограничивает количество запросов
type RateLimiter struct {
	tokens chan struct{}
	ticker *time.Ticker
}

// NewRateLimiter создает новый rate limiter
// rate - количество запросов в секунду
// burst - максимальное количество токенов
func NewRateLimiter(rate int, burst int) *RateLimiter {
	tokens := make(chan struct{}, burst)
	for i := 0; i < burst; i++ {
		tokens <- struct{}{}
	}

	rl := &RateLimiter{
		tokens: tokens,
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
			rl.tokens <- struct{}{}
			fmt.Println("add token")
		}
	}()
}

// Allow проверяет, можно ли выполнить запрос
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

// Wait ждет, пока не станет доступен токен
func (rl *RateLimiter) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-rl.tokens:
		return nil
	}

}

// Stop останавливает rate limiter
func (rl *RateLimiter) Stop() {
	if rl.ticker != nil {
		rl.ticker.Stop()
	}
	close(rl.tokens)
}
