package task9

import (
	"context"
	"testing"
	"time"
)

func TestRateLimiterBasic(t *testing.T) {
	rl := NewRateLimiter(10, 5) // 10 запросов в секунду, burst 5

	// Первые 5 запросов должны пройти сразу (burst)
	for i := 0; i < 5; i++ {
		if !rl.Allow() {
			t.Errorf("Request %d should be allowed (burst)", i+1)
		}
	}

	// 6-й запрос должен быть отклонен (burst исчерпан)
	if rl.Allow() {
		t.Error("6th request should be denied (burst exhausted)")
	}

	rl.Stop()
}

func TestRateLimiterRefill(t *testing.T) {
	rl := NewRateLimiter(2, 1) // 2 запроса в секунду, burst 1

	// Первый запрос должен пройти
	if !rl.Allow() {
		t.Error("First request should be allowed")
	}

	// Второй запрос должен быть отклонен
	if rl.Allow() {
		t.Error("Second request should be denied")
	}

	// Ждем пополнения токенов
	time.Sleep(600 * time.Millisecond)

	// Теперь запрос должен пройти
	if !rl.Allow() {
		t.Error("Request after refill should be allowed")
	}

	rl.Stop()
}

func TestRateLimiterWait(t *testing.T) {
	rl := NewRateLimiter(2, 1) // 2 запроса в секунду, burst 1

	// Первый запрос должен пройти
	if !rl.Allow() {
		t.Error("First request should be allowed")
	}

	// Второй запрос должен быть отклонен
	if rl.Allow() {
		t.Error("Second request should be denied")
	}

	// Тестируем Wait
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	start := time.Now()
	err := rl.Wait(ctx)
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Wait should succeed, got error: %v", err)
	}

	// Проверяем, что ждали примерно 500ms (половина секунды для 2 req/sec)
	if elapsed < 400*time.Millisecond || elapsed > 800*time.Millisecond {
		t.Errorf("Wait took %v, expected ~500ms", elapsed)
	}

	rl.Stop()
}

func TestRateLimiterWaitTimeout(t *testing.T) {
	rl := NewRateLimiter(1, 1) // 1 запрос в секунду, burst 1

	// Исчерпываем burst
	if !rl.Allow() {
		t.Error("First request should be allowed")
	}

	// Тестируем Wait с коротким timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := rl.Wait(ctx)
	if err == nil {
		t.Error("Wait should timeout")
	}

	rl.Stop()
}

func TestRateLimiterHighRate(t *testing.T) {
	rl := NewRateLimiter(100, 10) // 100 запросов в секунду, burst 10

	// Первые 10 запросов должны пройти
	allowed := 0
	for i := 0; i < 15; i++ {
		if rl.Allow() {
			allowed++
		}
	}

	if allowed != 10 {
		t.Errorf("Expected 10 allowed requests, got %d", allowed)
	}

	rl.Stop()
}

func TestRateLimiterZeroRate(t *testing.T) {
	rl := NewRateLimiter(0, 5) // 0 запросов в секунду, burst 5

	// Первые 5 запросов должны пройти (burst)
	for i := 0; i < 5; i++ {
		if !rl.Allow() {
			t.Errorf("Request %d should be allowed (burst)", i+1)
		}
	}

	// Дальнейшие запросы должны быть отклонены
	for i := 0; i < 5; i++ {
		if rl.Allow() {
			t.Error("Request after burst should be denied")
		}
	}

	rl.Stop()
}

func TestRateLimiterConcurrent(t *testing.T) {
	rl := NewRateLimiter(10, 5) // 10 запросов в секунду, burst 5

	// Запускаем несколько горутин, которые пытаются получить токены
	results := make(chan bool, 20)
	for i := 0; i < 20; i++ {
		go func() {
			results <- rl.Allow()
		}()
	}

	// Собираем результаты
	allowed := 0
	for i := 0; i < 20; i++ {
		if <-results {
			allowed++
		}
	}

	// Должно быть разрешено ровно 5 запросов (burst)
	if allowed != 5 {
		t.Errorf("Expected 5 allowed requests, got %d", allowed)
	}

	rl.Stop()
}

