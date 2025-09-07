package task11

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestSemaphoreBasic(t *testing.T) {
	sem := NewSemaphore(2)

	// Первые два запроса должны пройти
	if !sem.TryAcquire() {
		t.Error("First acquire should succeed")
	}
	if !sem.TryAcquire() {
		t.Error("Second acquire should succeed")
	}

	// Третий запрос должен быть отклонен
	if sem.TryAcquire() {
		t.Error("Third acquire should fail")
	}

	// Освобождаем один токен
	sem.Release()

	// Теперь запрос должен пройти
	if !sem.TryAcquire() {
		t.Error("Acquire after release should succeed")
	}

	// Снова должен быть отклонен
	if sem.TryAcquire() {
		t.Error("Acquire should fail when semaphore is full")
	}
}

func TestSemaphoreAcquire(t *testing.T) {
	sem := NewSemaphore(1)

	// Первый запрос должен пройти
	sem.Acquire()

	// Второй запрос должен блокироваться
	done := make(chan bool)
	go func() {
		sem.Acquire()
		done <- true
	}()

	// Проверяем, что горутина заблокирована
	select {
	case <-done:
		t.Error("Acquire should block when semaphore is empty")
	case <-time.After(100 * time.Millisecond):
		// Ожидаемое поведение
	}

	// Освобождаем токен
	sem.Release()

	// Теперь горутина должна разблокироваться
	select {
	case <-done:
		// Ожидаемое поведение
	case <-time.After(100 * time.Millisecond):
		t.Error("Acquire should succeed after release")
	}
}

func TestSemaphoreAcquireWithContext(t *testing.T) {
	sem := NewSemaphore(1)

	// Исчерпываем семафор
	sem.Acquire()

	// Тестируем AcquireWithContext с timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := sem.AcquireWithContext(ctx)
	elapsed := time.Since(start)

	if err == nil {
		t.Error("AcquireWithContext should timeout")
	}

	if elapsed < 90*time.Millisecond || elapsed > 150*time.Millisecond {
		t.Errorf("Timeout should be ~100ms, got %v", elapsed)
	}
}

func TestSemaphoreAcquireWithContextSuccess(t *testing.T) {
	sem := NewSemaphore(1)

	// Тестируем успешный AcquireWithContext
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := sem.AcquireWithContext(ctx)
	if err != nil {
		t.Errorf("AcquireWithContext should succeed, got error: %v", err)
	}
}

func TestSemaphoreAcquireWithContextCancelled(t *testing.T) {
	sem := NewSemaphore(1)

	// Исчерпываем семафор
	sem.Acquire()

	// Тестируем отмену контекста
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	err := sem.AcquireWithContext(ctx)
	if err == nil {
		t.Error("AcquireWithContext should be cancelled")
	}
}

func TestSemaphoreExample(t *testing.T) {
	results := SemaphoreExample(5, 2) // 5 воркеров, максимум 2 одновременно

	if len(results) != 5 {
		t.Errorf("Expected 5 results, got %d", len(results))
	}

	// Проверяем, что все результаты содержат "completed"
	for i, result := range results {
		if result != "worker completed" {
			t.Errorf("Result %d should be 'worker completed', got '%s'", i, result)
		}
	}
}

func TestSemaphoreConcurrent(t *testing.T) {
	sem := NewSemaphore(3)
	var wg sync.WaitGroup
	results := make(chan bool, 10)

	// Запускаем 10 горутин, которые пытаются получить токены
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- sem.TryAcquire()
		}()
	}

	wg.Wait()
	close(results)

	// Считаем успешные запросы
	successful := 0
	for success := range results {
		if success {
			successful++
		}
	}

	// Должно быть ровно 3 успешных запроса
	if successful != 3 {
		t.Errorf("Expected 3 successful acquires, got %d", successful)
	}
}

func TestSemaphoreZeroCapacity(t *testing.T) {
	sem := NewSemaphore(0)

	// Все запросы должны быть отклонены
	if sem.TryAcquire() {
		t.Error("TryAcquire should fail with zero capacity")
	}

	// Acquire должен блокироваться навсегда
	done := make(chan bool)
	go func() {
		sem.Acquire()
		done <- true
	}()

	select {
	case <-done:
		t.Error("Acquire should block forever with zero capacity")
	case <-time.After(100 * time.Millisecond):
		// Ожидаемое поведение
	}
}

func TestSemaphoreReleaseWithoutAcquire(t *testing.T) {
	sem := NewSemaphore(1)

	// Освобождаем токен без предварительного получения
	// Это не должно вызывать панику
	sem.Release()

	// Теперь должно быть доступно 2 токена
	if !sem.TryAcquire() {
		t.Error("First acquire should succeed")
	}
	if !sem.TryAcquire() {
		t.Error("Second acquire should succeed")
	}
}

