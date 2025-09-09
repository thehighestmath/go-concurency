package task13

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestCircuitBreakerBasic(t *testing.T) {
	config := CircuitBreakerConfig{
		FailureThreshold: 3,
		SuccessThreshold: 2,
		Timeout:          100,
		MaxRequests:      1,
	}

	cb := NewCircuitBreaker(config)

	// Тестируем успешные вызовы
	for i := 0; i < 2; i++ {
		result, err := cb.Execute(func() (any, error) {
			return "success", nil
		})

		if err != nil {
			t.Errorf("Expected success, got error: %v", err)
		}
		if result != "success" {
			t.Errorf("Expected 'success', got %v", result)
		}
	}

	// Проверяем состояние
	if cb.GetState() != StateClosed {
		t.Errorf("Expected StateClosed, got %v", cb.GetState())
	}
}

func TestCircuitBreakerFailureThreshold(t *testing.T) {
	config := CircuitBreakerConfig{
		FailureThreshold: 2,
		SuccessThreshold: 1,
		Timeout:          100,
		MaxRequests:      1,
	}

	cb := NewCircuitBreaker(config)

	// Вызываем функцию, которая всегда падает
	for i := 0; i < 2; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			return nil, &TestError{message: "test error"}
		})

		if err == nil {
			t.Error("Expected error, got success")
		}
	}

	// Проверяем, что circuit breaker открыт
	if cb.GetState() != StateOpen {
		t.Errorf("Expected StateOpen, got %v", cb.GetState())
	}
}

func TestCircuitBreakerOpenState(t *testing.T) {
	config := CircuitBreakerConfig{
		FailureThreshold: 1,
		SuccessThreshold: 1,
		Timeout:          100,
		MaxRequests:      1,
	}

	cb := NewCircuitBreaker(config)

	// Открываем circuit breaker
	cb.Execute(func() (interface{}, error) {
		return nil, &TestError{message: "test error"}
	})

	// Проверяем, что вызовы блокируются
	_, err := cb.Execute(func() (interface{}, error) {
		return "success", nil
	})

	if err == nil {
		t.Error("Expected error in open state, got success")
	}
}

func TestCircuitBreakerHalfOpenState(t *testing.T) {
	config := CircuitBreakerConfig{
		FailureThreshold: 1,
		SuccessThreshold: 1,
		Timeout:          50,
		MaxRequests:      1,
	}

	cb := NewCircuitBreaker(config)

	// Открываем circuit breaker
	cb.Execute(func() (interface{}, error) {
		return nil, &TestError{message: "test error"}
	})

	// Ждем перехода в half-open
	time.Sleep(60 * time.Millisecond)

	// Проверяем состояние
	if cb.GetState() != StateHalfOpen {
		t.Errorf("Expected StateHalfOpen, got %v", cb.GetState())
	}
}

func TestCircuitBreakerRecovery(t *testing.T) {
	config := CircuitBreakerConfig{
		FailureThreshold: 1,
		SuccessThreshold: 1,
		Timeout:          50,
		MaxRequests:      1,
	}

	cb := NewCircuitBreaker(config)

	// Открываем circuit breaker
	cb.Execute(func() (interface{}, error) {
		return nil, &TestError{message: "test error"}
	})

	// Ждем перехода в half-open
	time.Sleep(60 * time.Millisecond)

	// Успешный вызов должен закрыть circuit breaker
	result, err := cb.Execute(func() (interface{}, error) {
		return "success", nil
	})

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if result != "success" {
		t.Errorf("Expected 'success', got %v", result)
	}

	// Проверяем, что circuit breaker закрыт
	if cb.GetState() != StateClosed {
		t.Errorf("Expected StateClosed, got %v", cb.GetState())
	}
}

func TestCircuitBreakerStats(t *testing.T) {
	config := CircuitBreakerConfig{
		FailureThreshold: 3,
		SuccessThreshold: 2,
		Timeout:          100,
		MaxRequests:      1,
	}

	cb := NewCircuitBreaker(config)

	// Выполняем несколько операций
	cb.Execute(func() (interface{}, error) {
		return "success", nil
	})

	cb.Execute(func() (interface{}, error) {
		return nil, &TestError{message: "test error"}
	})

	stats := cb.GetStats()

	if stats["total"] != 2 {
		t.Errorf("Expected total 2, got %d", stats["total"])
	}
	if stats["success"] != 1 {
		t.Errorf("Expected success 1, got %d", stats["success"])
	}
	if stats["failure"] != 1 {
		t.Errorf("Expected failure 1, got %d", stats["failure"])
	}
}

// TestError для тестирования
type TestError struct {
	message string
}

func (e *TestError) Error() string {
	return e.message
}

func successfulCall() (any, error) {
	return "ok", nil
}
func failedCall() (any, error) {
	return nil, errors.New("fail")
}

func TestCircuitBreaker_StartsClosed_AllowsCalls(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 2, SuccessThreshold: 1, Timeout: 100, MaxRequests: 1})
	_, err := cb.Execute(successfulCall)
	if err != nil {
		t.Errorf("Expected call to be allowed in Closed, got error: %v", err)
	}
}

func TestCircuitBreaker_TripsToOpenAfterFailures(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 2, SuccessThreshold: 1, Timeout: 100, MaxRequests: 1})

	_, _ = cb.Execute(failedCall)
	_, _ = cb.Execute(failedCall)

	if cb.GetState() != StateOpen {
		t.Errorf("Expected breaker to be Open after failures, got: %v", cb.GetState())
	}
	_, err := cb.Execute(successfulCall)
	if err == nil {
		t.Errorf("Expected call to fail while Open")
	}
}

func TestCircuitBreaker_TransitionsToHalfOpenAfterTimeout(t *testing.T) {
	timeout := int64(100)
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 1, SuccessThreshold: 1, Timeout: timeout, MaxRequests: 1})
	_, _ = cb.Execute(failedCall) // Trip to Open

	// Wait for timeout + a little extra
	time.Sleep(time.Duration(timeout+20) * time.Millisecond)
	state := cb.GetState()
	if state != StateHalfOpen {
		t.Errorf("Expected state HalfOpen after timeout, got: %v", state)
	}
}

func TestCircuitBreaker_FromHalfOpen_ReturnsToClosedOnSuccesses(t *testing.T) {
	timeout := int64(50)
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 1, SuccessThreshold: 2, Timeout: timeout, MaxRequests: 1})
	_, _ = cb.Execute(failedCall) // Go to Open
	time.Sleep(time.Duration(timeout+10) * time.Millisecond)
	_, _ = cb.Execute(successfulCall)
	_, _ = cb.Execute(successfulCall)

	if cb.GetState() != StateClosed {
		t.Errorf("Expected state Closed after successes in HalfOpen, got: %v", cb.GetState())
	}
}

func TestCircuitBreaker_FromHalfOpen_ReturnsToOpenOnFailure(t *testing.T) {
	timeout := int64(50)
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 1, SuccessThreshold: 2, Timeout: timeout, MaxRequests: 1})
	_, _ = cb.Execute(failedCall) // Go to Open
	time.Sleep(time.Duration(timeout+10) * time.Millisecond)
	_, _ = cb.Execute(failedCall)

	if cb.GetState() != StateOpen {
		t.Errorf("Expected state Open after failure in HalfOpen, got: %v", cb.GetState())
	}
}

func TestCircuitBreaker_Stats(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 3, SuccessThreshold: 1, Timeout: 50, MaxRequests: 1})
	_, _ = cb.Execute(successfulCall)
	_, _ = cb.Execute(failedCall)
	_, _ = cb.Execute(failedCall)
	stats := cb.GetStats()
	expectedTotal := cb.SuccessCount + cb.FailureCount
	if stats["success"] != cb.SuccessCount || stats["failure"] != cb.FailureCount || stats["total"] != expectedTotal {
		t.Errorf("Unexpected stats: %v", stats)
	}
}

func TestCircuitBreaker_ConcurrentAccess(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 10, SuccessThreshold: 1, Timeout: 50, MaxRequests: 3})
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(j int) {
			if j%2 == 0 {
				cb.Execute(successfulCall)
			} else {
				cb.Execute(failedCall)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	// State may or may not be Open depending on the order, but the test ensures concurrent safety
}

func successCall() (any, error) { return "ok", nil }
func failCall() (any, error)    { return nil, errors.New("fail") }

func TestBreaker_DoesNotTripBeforeThreshold(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 3, SuccessThreshold: 2, Timeout: 50, MaxRequests: 2})

	_, _ = cb.Execute(failCall)
	_, _ = cb.Execute(successCall)
	if cb.GetState() != StateClosed {
		t.Errorf("Breaker opened prematurely; expected Closed, got %v", cb.GetState())
	}
}

func TestBreaker_OpenRemainsBlocked(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 1, SuccessThreshold: 1, Timeout: 100, MaxRequests: 1})

	_, _ = cb.Execute(failCall)
	for i := 0; i < 3; i++ {
		out, err := cb.Execute(successCall)
		if err == nil {
			t.Errorf("Should block in Open, got result=%v", out)
		}
	}
}

func TestBreaker_HalfOpen_AllowsLimited(t *testing.T) {
	cfg := CircuitBreakerConfig{FailureThreshold: 1, SuccessThreshold: 2, Timeout: 30, MaxRequests: 2}
	cb := NewCircuitBreaker(cfg)
	_, _ = cb.Execute(failCall)

	time.Sleep(35 * time.Millisecond)
	state := cb.GetState()
	if state != StateHalfOpen {
		t.Fatalf("Expected half-open after timeout, got %v", state)
	}

	_, err1 := cb.Execute(successCall)
	_, err2 := cb.Execute(successCall)
	if err1 != nil || err2 != nil {
		t.Errorf("Should allow MaxRequests in half-open, got err1=%v err2=%v", err1, err2)
	}
}

func TestBreaker_Counters_ResetCorrectly(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 2, SuccessThreshold: 2, Timeout: 20, MaxRequests: 2})
	_, _ = cb.Execute(failCall)
	_, _ = cb.Execute(failCall)
	time.Sleep(25 * time.Millisecond)
	cb.GetState()
	// In half-open: success resets FailureCount
	_, _ = cb.Execute(successCall)
	if cb.FailureCount != 0 {
		t.Errorf("FailureCount should be reset on success, got %d", cb.FailureCount)
	}
}

func TestBreaker_StatsReport(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 2, SuccessThreshold: 2, Timeout: 50, MaxRequests: 2})
	_, _ = cb.Execute(failCall)
	_, _ = cb.Execute(successCall)
	stats := cb.GetStats()
	if stats["success"] != 1 || stats["failure"] != 1 || stats["total"] != 2 {
		t.Errorf("Stats not reported correctly: %v", stats)
	}
}

func TestBreaker_ConcurrentStress(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{FailureThreshold: 3, SuccessThreshold: 2, Timeout: 20, MaxRequests: 1})
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n%4 == 0 {
				cb.Execute(failCall)
			} else {
				cb.Execute(successCall)
			}
		}(i)
	}
	wg.Wait()
	// Just ensure no panic/race
}
