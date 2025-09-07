package task15

import (
	"testing"
	"time"
)

func TestRetryManagerBasic(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   3,
		BaseDelay:     10,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	attempts := 0
	err := rm.Execute(nil, func() error {
		attempts++
		if attempts < 3 {
			return &TestError{message: "temporary error"}
		}
		return nil
	})

	if err != nil {
		t.Errorf("Expected success after retries, got error: %v", err)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestRetryManagerMaxAttempts(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   2,
		BaseDelay:     10,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	attempts := 0
	err := rm.Execute(nil, func() error {
		attempts++
		return &TestError{message: "persistent error"}
	})

	if err == nil {
		t.Error("Expected error after max attempts, got success")
	}

	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
}

func TestRetryManagerExponentialBackoff(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   3,
		BaseDelay:     10,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	start := time.Now()
	attempts := 0

	err := rm.Execute(nil, func() error {
		attempts++
		if attempts < 3 {
			return &TestError{message: "temporary error"}
		}
		return nil
	})

	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}

	// Проверяем, что задержки увеличиваются экспоненциально
	// Первая задержка: 10ms, вторая: 20ms, общее время: ~30ms
	if elapsed < 25*time.Millisecond || elapsed > 50*time.Millisecond {
		t.Errorf("Expected exponential backoff, elapsed: %v", elapsed)
	}
}

func TestRetryManagerMaxDelay(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   5,
		BaseDelay:     50,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	start := time.Now()
	attempts := 0

	err := rm.Execute(nil, func() error {
		attempts++
		if attempts < 5 {
			return &TestError{message: "temporary error"}
		}
		return nil
	})

	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}

	// Проверяем, что задержки не превышают MaxDelay
	// 50ms + 100ms + 100ms + 100ms = ~350ms
	if elapsed < 300*time.Millisecond || elapsed > 400*time.Millisecond {
		t.Errorf("Expected max delay limit, elapsed: %v", elapsed)
	}
}

func TestRetryManagerJitter(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   3,
		BaseDelay:     50,
		MaxDelay:      200,
		BackoffFactor: 2.0,
		Jitter:        true,
	}

	rm := NewRetryManager(config)

	start := time.Now()
	attempts := 0

	err := rm.Execute(nil, func() error {
		attempts++
		if attempts < 3 {
			return &TestError{message: "temporary error"}
		}
		return nil
	})

	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}

	// С jitter время должно быть немного случайным
	if elapsed < 50*time.Millisecond || elapsed > 200*time.Millisecond {
		t.Errorf("Expected jittered delay, elapsed: %v", elapsed)
	}
}

func TestRetryManagerCustomRetry(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   5,
		BaseDelay:     10,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	attempts := 0
	err := rm.ExecuteWithCustomRetry(nil, func() error {
		attempts++
		return &TestError{message: "custom error"}
	}, func(err error) bool {
		// Повторяем только первые 2 раза
		return attempts < 2
	})

	if err == nil {
		t.Error("Expected error after custom retry limit, got success")
	}

	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
}

func TestRetryManagerStats(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   3,
		BaseDelay:     10,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	// Выполняем успешную операцию
	rm.Execute(nil, func() error {
		return nil
	})

	// Выполняем операцию с ошибкой
	rm.Execute(nil, func() error {
		return &TestError{message: "error"}
	})

	stats := rm.GetStats()

	if stats["attempts"] != 2 {
		t.Errorf("Expected 2 attempts, got %d", stats["attempts"])
	}
	if stats["successes"] != 1 {
		t.Errorf("Expected 1 success, got %d", stats["successes"])
	}
	if stats["failures"] != 1 {
		t.Errorf("Expected 1 failure, got %d", stats["failures"])
	}
}

func TestRetryManagerResetStats(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   3,
		BaseDelay:     10,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	// Выполняем операцию
	rm.Execute(nil, func() error {
		return nil
	})

	// Сбрасываем статистику
	rm.ResetStats()

	stats := rm.GetStats()

	if stats["attempts"] != 0 {
		t.Errorf("Expected 0 attempts after reset, got %d", stats["attempts"])
	}
}

func TestRetryManagerZeroAttempts(t *testing.T) {
	config := RetryConfig{
		MaxAttempts:   0,
		BaseDelay:     10,
		MaxDelay:      100,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	rm := NewRetryManager(config)

	attempts := 0
	err := rm.Execute(nil, func() error {
		attempts++
		return &TestError{message: "error"}
	})

	if err == nil {
		t.Error("Expected error with zero attempts, got success")
	}

	if attempts != 0 {
		t.Errorf("Expected 0 attempts, got %d", attempts)
	}
}
