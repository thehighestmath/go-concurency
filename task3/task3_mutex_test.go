package task3

import (
	"sync"
	"testing"
)

func TestSafeCounterBasic(t *testing.T) {
	counter := NewSafeCounter()

	// Проверяем начальное значение
	if counter.GetValue() != 0 {
		t.Errorf("Expected initial value 0, got %d", counter.GetValue())
	}

	// Тестируем Increment
	counter.Increment()
	if counter.GetValue() != 1 {
		t.Errorf("Expected value 1 after increment, got %d", counter.GetValue())
	}

	counter.Increment()
	if counter.GetValue() != 2 {
		t.Errorf("Expected value 2 after second increment, got %d", counter.GetValue())
	}

	// Тестируем Decrement
	counter.Decrement()
	if counter.GetValue() != 1 {
		t.Errorf("Expected value 1 after decrement, got %d", counter.GetValue())
	}

	counter.Decrement()
	if counter.GetValue() != 0 {
		t.Errorf("Expected value 0 after second decrement, got %d", counter.GetValue())
	}

	// Тестируем Decrement ниже нуля
	counter.Decrement()
	if counter.GetValue() != -1 {
		t.Errorf("Expected value -1 after decrement below zero, got %d", counter.GetValue())
	}
}

func TestSafeCounterConcurrent(t *testing.T) {
	counter := NewSafeCounter()
	const numGoroutines = 100
	const operationsPerGoroutine = 100

	var wg sync.WaitGroup

	// Запускаем горутины, которые увеличивают счетчик
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	expected := numGoroutines * operationsPerGoroutine
	if counter.GetValue() != expected {
		t.Errorf("Expected value %d after concurrent increments, got %d", expected, counter.GetValue())
	}
}

func TestSafeCounterMixedOperations(t *testing.T) {
	counter := NewSafeCounter()
	const numGoroutines = 50

	var wg sync.WaitGroup

	// Половина горутин увеличивает, половина уменьшает
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(increment bool) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				if increment {
					counter.Increment()
				} else {
					counter.Decrement()
				}
			}
		}(i%2 == 0)
	}

	wg.Wait()

	// Ожидаемое значение: 25 горутин * 100 операций * (+1) + 25 горутин * 100 операций * (-1) = 0
	expected := 0
	if counter.GetValue() != expected {
		t.Errorf("Expected value %d after mixed operations, got %d", expected, counter.GetValue())
	}
}

func TestSafeCounterMultipleInstances(t *testing.T) {
	counter1 := NewSafeCounter()
	counter2 := NewSafeCounter()

	counter1.Increment()
	counter1.Increment()

	counter2.Decrement()

	if counter1.GetValue() != 2 {
		t.Errorf("Counter1 expected 2, got %d", counter1.GetValue())
	}

	if counter2.GetValue() != -1 {
		t.Errorf("Counter2 expected -1, got %d", counter2.GetValue())
	}
}
