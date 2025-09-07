package task13

import (
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

