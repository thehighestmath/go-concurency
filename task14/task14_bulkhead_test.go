package task14

import (
	"sync"
	"testing"
	"time"
)

func TestBulkheadManagerBasic(t *testing.T) {
	bm := NewBulkheadManager()

	// Настраиваем пулы
	bm.ConfigurePool(OperationRead, 2)
	bm.ConfigurePool(OperationWrite, 1)
	bm.ConfigurePool(OperationDelete, 1)

	// Тестируем выполнение операций
	var wg sync.WaitGroup
	results := make([]bool, 3)

	// Запускаем операции чтения
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			err := bm.Execute(OperationRead, func() error {
				time.Sleep(10 * time.Millisecond)
				return nil
			})
			results[index] = (err == nil)
		}(i)
	}

	// Запускаем операцию записи
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := bm.Execute(OperationWrite, func() error {
			time.Sleep(10 * time.Millisecond)
			return nil
		})
		results[2] = (err == nil)
	}()

	wg.Wait()

	// Проверяем результаты
	for i, result := range results {
		if !result {
			t.Errorf("Operation %d should succeed", i)
		}
	}
}

func TestBulkheadManagerConcurrencyLimit(t *testing.T) {
	bm := NewBulkheadManager()
	bm.ConfigurePool(OperationRead, 1) // Только 1 одновременная операция

	var wg sync.WaitGroup
	results := make([]bool, 3)

	// Запускаем 3 операции чтения
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			err := bm.Execute(OperationRead, func() error {
				time.Sleep(50 * time.Millisecond)
				return nil
			})
			results[index] = (err == nil)
		}(i)
	}

	wg.Wait()

	// Все операции должны выполниться успешно
	// (bulkhead не блокирует, а ограничивает конкурентность)
	successCount := 0
	for _, result := range results {
		if result {
			successCount++
		}
	}

	if successCount != 3 {
		t.Errorf("Expected 3 successful operations, got %d", successCount)
	}
}

func TestBulkheadManagerDifferentOperations(t *testing.T) {
	bm := NewBulkheadManager()
	bm.ConfigurePool(OperationRead, 2)
	bm.ConfigurePool(OperationWrite, 1)
	bm.ConfigurePool(OperationDelete, 1)

	var wg sync.WaitGroup
	results := make(map[OperationType]bool)

	// Тестируем разные типы операций
	operations := []OperationType{OperationRead, OperationWrite, OperationDelete}

	for _, opType := range operations {
		wg.Add(1)
		go func(op OperationType) {
			defer wg.Done()
			err := bm.Execute(op, func() error {
				time.Sleep(10 * time.Millisecond)
				return nil
			})
			results[op] = (err == nil)
		}(opType)
	}

	wg.Wait()

	// Все операции должны выполниться успешно
	for _, opType := range operations {
		if !results[opType] {
			t.Errorf("Operation %v should succeed", opType)
		}
	}
}

func TestBulkheadManagerStats(t *testing.T) {
	bm := NewBulkheadManager()
	bm.ConfigurePool(OperationRead, 2)
	bm.ConfigurePool(OperationWrite, 1)

	// Выполняем операции
	bm.Execute(OperationRead, func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	})

	bm.Execute(OperationWrite, func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	})

	// Проверяем статистику
	readStats := bm.GetPoolStats(OperationRead)
	writeStats := bm.GetPoolStats(OperationWrite)

	if readStats == nil || writeStats == nil {
		t.Error("Stats should not be nil")
	}

	totalStats := bm.GetTotalStats()
	if totalStats == nil {
		t.Error("Total stats should not be nil")
	}
}

func TestBulkheadManagerIsolation(t *testing.T) {
	bm := NewBulkheadManager()
	bm.ConfigurePool(OperationRead, 2)
	bm.ConfigurePool(OperationWrite, 1)

	var wg sync.WaitGroup
	start := time.Now()

	// Запускаем операцию записи (должна выполниться быстро)
	wg.Add(1)
	go func() {
		defer wg.Done()
		bm.Execute(OperationWrite, func() error {
			time.Sleep(10 * time.Millisecond)
			return nil
		})
	}()

	// Запускаем много операций чтения
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bm.Execute(OperationRead, func() error {
				time.Sleep(50 * time.Millisecond)
				return nil
			})
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)

	// Операция записи не должна блокироваться операциями чтения
	if elapsed > 200*time.Millisecond {
		t.Errorf("Operations took too long: %v (bulkhead should provide isolation)", elapsed)
	}
}

func TestBulkheadManagerZeroConcurrency(t *testing.T) {
	bm := NewBulkheadManager()
	bm.ConfigurePool(OperationRead, 0) // Нулевая конкурентность

	// Операция должна завершиться (не блокироваться навсегда)
	done := make(chan bool)
	go func() {
		err := bm.Execute(OperationRead, func() error {
			return nil
		})
		done <- (err == nil)
	}()

	select {
	case result := <-done:
		if !result {
			t.Error("Operation should succeed even with zero concurrency")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Operation should not block forever")
	}
}

