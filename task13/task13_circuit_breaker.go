package task13

import (
	"fmt"
	"sync"
	"time"
)

// Task 13: Circuit Breaker Pattern
// Реализуйте паттерн Circuit Breaker для защиты от каскадных сбоев.
// Circuit Breaker должен переключаться между состояниями: Closed, Open, Half-Open.

// CircuitState представляет состояние circuit breaker
type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

// CircuitBreaker реализует паттерн Circuit Breaker
type CircuitBreaker struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - текущее состояние
	// - счетчики успешных/неуспешных вызовов
	// - настройки (threshold, timeout, etc.)
	// - mutex для синхронизации
	// - время последнего сбоя
	Config            CircuitBreakerConfig
	State             CircuitState
	SuccessCount      int
	FailureCount      int
	TotalFailureCount int
	Mu                sync.Mutex
	LastTime          time.Time
}

// CircuitBreakerConfig содержит настройки circuit breaker
type CircuitBreakerConfig struct {
	FailureThreshold int   // Количество сбоев для перехода в Open
	SuccessThreshold int   // Количество успехов для перехода в Closed
	Timeout          int64 // Время в Open состоянии (мс)
	MaxRequests      int   // Максимум запросов в Half-Open состоянии
}

// NewCircuitBreaker создает новый circuit breaker
func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker {
	// TODO: Реализуйте конструктор
	return &CircuitBreaker{
		Config:            config,
		State:             StateClosed,
		SuccessCount:      0,
		FailureCount:      0,
		TotalFailureCount: 0,
		Mu:                sync.Mutex{},
		LastTime:          time.Time{},
	}
}

// Execute выполняет функцию через circuit breaker
func (cb *CircuitBreaker) Execute(fn func() (any, error)) (any, error) {
	cb.GetState()
	if cb.State == StateOpen {
		return nil, fmt.Errorf("qwe")
	}

	res, err := fn()

	if err != nil {
		cb.LastTime = time.Now()
		cb.TotalFailureCount++
		cb.FailureCount++
		if cb.State == StateHalfOpen {
			cb.State = StateOpen
		}
	} else {
		cb.FailureCount = 0
		cb.SuccessCount++
		if cb.State == StateHalfOpen {
			cb.State = StateClosed
		}
	}
	if cb.FailureCount == cb.Config.FailureThreshold {
		cb.State = StateOpen
	}

	return res, err
}

// GetState возвращает текущее состояние circuit breaker
func (cb *CircuitBreaker) GetState() CircuitState {
	timePassed := cb.LastTime.Add(time.Duration(cb.Config.Timeout) * time.Millisecond)
	if cb.State == StateOpen && time.Now().After(timePassed) {
		cb.State = StateHalfOpen
	}
	return cb.State
}

// GetStats возвращает статистику circuit breaker
func (cb *CircuitBreaker) GetStats() map[string]int {
	return map[string]int{
		"success": cb.SuccessCount,
		"failure": cb.TotalFailureCount,
		"total":   cb.TotalFailureCount + cb.SuccessCount,
	}
}
