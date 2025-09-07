package task5

import (
	"testing"
	"time"
)

func TestWorkerPoolBasic(t *testing.T) {
	pool := NewWorkerPool(2)
	pool.Start()

	// Отправляем задачи
	task1 := Task{ID: 1, Data: "task1"}
	task2 := Task{ID: 2, Data: "task2"}

	pool.SubmitTask(task1)
	pool.SubmitTask(task2)

	// Собираем результаты
	results := make([]Result, 0)
	resultChan := pool.GetResult()

	timeout := time.After(1 * time.Second)
	for i := 0; i < 2; i++ {
		select {
		case result := <-resultChan:
			results = append(results, result)
		case <-timeout:
			t.Fatal("Timeout waiting for results")
		}
	}

	// Проверяем результаты
	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}

	// Проверяем, что все задачи обработаны
	taskIDs := make(map[int]bool)
	for _, result := range results {
		taskIDs[result.TaskID] = true
		if result.Error != nil {
			t.Errorf("Task %d failed: %v", result.TaskID, result.Error)
		}
	}

	if !taskIDs[1] || !taskIDs[2] {
		t.Error("Not all tasks were processed")
	}

	pool.Stop()
}

func TestWorkerPoolManyTasks(t *testing.T) {
	pool := NewWorkerPool(3)
	pool.Start()

	// Отправляем много задач
	numTasks := 10
	for i := 0; i < numTasks; i++ {
		task := Task{ID: i, Data: "task"}
		pool.SubmitTask(task)
	}

	// Собираем результаты
	results := make([]Result, 0)
	resultChan := pool.GetResult()

	timeout := time.After(2 * time.Second)
	for i := 0; i < numTasks; i++ {
		select {
		case result := <-resultChan:
			results = append(results, result)
		case <-timeout:
			t.Fatalf("Timeout waiting for result %d", i)
		}
	}

	if len(results) != numTasks {
		t.Errorf("Expected %d results, got %d", numTasks, len(results))
	}

	pool.Stop()
}

func TestWorkerPoolSingleWorker(t *testing.T) {
	pool := NewWorkerPool(1)
	pool.Start()

	task := Task{ID: 1, Data: "single task"}
	pool.SubmitTask(task)

	resultChan := pool.GetResult()
	timeout := time.After(1 * time.Second)

	select {
	case result := <-resultChan:
		if result.TaskID != 1 {
			t.Errorf("Expected task ID 1, got %d", result.TaskID)
		}
		if result.Error != nil {
			t.Errorf("Task failed: %v", result.Error)
		}
	case <-timeout:
		t.Fatal("Timeout waiting for result")
	}

	pool.Stop()
}

func TestWorkerPoolNoTasks(t *testing.T) {
	pool := NewWorkerPool(2)
	pool.Start()

	// Не отправляем задачи
	resultChan := pool.GetResult()

	// Проверяем, что канал результатов пуст
	select {
	case result := <-resultChan:
		t.Errorf("Unexpected result: %v", result)
	default:
		// Ожидаемое поведение - канал пуст
	}

	pool.Stop()
}

