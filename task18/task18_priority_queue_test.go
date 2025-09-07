package task18

import (
	"sync"
	"testing"
	"time"
)

func TestPriorityQueueBasic(t *testing.T) {
	pq := NewPriorityQueue()

	// Добавляем элементы с разными приоритетами
	pq.Push("low", 1)
	pq.Push("high", 10)
	pq.Push("medium", 5)

	// Проверяем размер
	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	// Проверяем, что очередь не пуста
	if pq.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	// Извлекаем элементы (должны быть в порядке приоритета)
	value, ok := pq.Pop()
	if !ok {
		t.Error("Should be able to pop element")
	}
	if value != "high" {
		t.Errorf("Expected 'high', got %v", value)
	}

	value, ok = pq.Pop()
	if !ok {
		t.Error("Should be able to pop element")
	}
	if value != "medium" {
		t.Errorf("Expected 'medium', got %v", value)
	}

	value, ok = pq.Pop()
	if !ok {
		t.Error("Should be able to pop element")
	}
	if value != "low" {
		t.Errorf("Expected 'low', got %v", value)
	}

	// Проверяем, что очередь пуста
	if !pq.IsEmpty() {
		t.Error("Queue should be empty")
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	pq := NewPriorityQueue()

	// Добавляем элементы
	pq.Push("low", 1)
	pq.Push("high", 10)

	// Проверяем peek
	value, ok := pq.Peek()
	if !ok {
		t.Error("Should be able to peek element")
	}
	if value != "high" {
		t.Errorf("Expected 'high', got %v", value)
	}

	// Проверяем, что элемент не извлечен
	if pq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", pq.Size())
	}
}

func TestPriorityQueueEmpty(t *testing.T) {
	pq := NewPriorityQueue()

	// Проверяем пустую очередь
	if !pq.IsEmpty() {
		t.Error("Queue should be empty")
	}

	value, ok := pq.Pop()
	if ok {
		t.Errorf("Should not be able to pop from empty queue, got %v", value)
	}

	value, ok = pq.Peek()
	if ok {
		t.Errorf("Should not be able to peek empty queue, got %v", value)
	}
}

func TestPriorityQueuePopBlocking(t *testing.T) {
	pq := NewPriorityQueue()

	// Запускаем горутину, которая ждет элемент
	done := make(chan bool)
	go func() {
		value := pq.PopBlocking()
		if value != "test" {
			t.Errorf("Expected 'test', got %v", value)
		}
		done <- true
	}()

	// Ждем немного, чтобы горутина запустилась
	time.Sleep(10 * time.Millisecond)

	// Добавляем элемент
	pq.Push("test", 5)

	// Проверяем, что горутина завершилась
	select {
	case <-done:
		// Ожидаемое поведение
	case <-time.After(100 * time.Millisecond):
		t.Error("PopBlocking should return when element is added")
	}
}

func TestPriorityQueueClear(t *testing.T) {
	pq := NewPriorityQueue()

	// Добавляем элементы
	pq.Push("item1", 1)
	pq.Push("item2", 2)

	// Очищаем очередь
	pq.Clear()

	// Проверяем, что очередь пуста
	if !pq.IsEmpty() {
		t.Error("Queue should be empty after clear")
	}

	if pq.Size() != 0 {
		t.Errorf("Expected size 0, got %d", pq.Size())
	}
}

func TestPriorityQueueSamePriority(t *testing.T) {
	pq := NewPriorityQueue()

	// Добавляем элементы с одинаковым приоритетом
	pq.Push("first", 5)
	pq.Push("second", 5)
	pq.Push("third", 5)

	// Извлекаем элементы
	values := make([]interface{}, 0, 3)
	for i := 0; i < 3; i++ {
		value, ok := pq.Pop()
		if !ok {
			t.Error("Should be able to pop element")
		}
		values = append(values, value)
	}

	// Проверяем, что все элементы извлечены
	if len(values) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(values))
	}
}

func TestConcurrentPriorityQueueBasic(t *testing.T) {
	cpq := NewConcurrentPriorityQueue(2)

	// Добавляем элементы
	cpq.Push("low", 1)
	cpq.Push("high", 10)
	cpq.Push("medium", 5)

	// Извлекаем элементы
	value, ok := cpq.Pop()
	if !ok {
		t.Error("Should be able to pop element")
	}
	if value != "high" {
		t.Errorf("Expected 'high', got %v", value)
	}
}

func TestConcurrentPriorityQueueConcurrent(t *testing.T) {
	cpq := NewConcurrentPriorityQueue(3)

	var wg sync.WaitGroup

	// Запускаем горутины, которые добавляют элементы
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			cpq.Push("item"+string(rune(index)), index)
		}(i)
	}

	wg.Wait()

	// Извлекаем все элементы
	values := make([]interface{}, 0, 10)
	for i := 0; i < 10; i++ {
		value, ok := cpq.Pop()
		if !ok {
			t.Error("Should be able to pop element")
		}
		values = append(values, value)
	}

	// Проверяем, что все элементы извлечены
	if len(values) != 10 {
		t.Errorf("Expected 10 elements, got %d", len(values))
	}
}

func TestConcurrentPriorityQueueStats(t *testing.T) {
	cpq := NewConcurrentPriorityQueue(2)

	// Добавляем элементы
	cpq.Push("item1", 1)
	cpq.Push("item2", 2)

	// Проверяем статистику
	stats := cpq.GetStats()
	if stats == nil {
		t.Error("Stats should not be nil")
	}
}

func TestPriorityQueueNegativePriority(t *testing.T) {
	pq := NewPriorityQueue()

	// Добавляем элементы с отрицательным приоритетом
	pq.Push("negative", -5)
	pq.Push("positive", 5)
	pq.Push("zero", 0)

	// Извлекаем элементы (должны быть в порядке приоритета)
	value, ok := pq.Pop()
	if !ok {
		t.Error("Should be able to pop element")
	}
	if value != "positive" {
		t.Errorf("Expected 'positive', got %v", value)
	}

	value, ok = pq.Pop()
	if !ok {
		t.Error("Should be able to pop element")
	}
	if value != "zero" {
		t.Errorf("Expected 'zero', got %v", value)
	}

	value, ok = pq.Pop()
	if !ok {
		t.Error("Should be able to pop element")
	}
	if value != "negative" {
		t.Errorf("Expected 'negative', got %v", value)
	}
}

