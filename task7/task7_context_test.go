package task7

import (
	"strings"
	"testing"
)

func TestProcessTasksWithContext(t *testing.T) {
	tasks := []LongRunningTask{
		{ID: 1, Duration: 50},
		{ID: 2, Duration: 30},
		{ID: 3, Duration: 40},
	}

	timeout := int64(200)
	results := ProcessTasksWithContext(tasks, timeout)

	if len(results) != len(tasks) {
		t.Errorf("Expected %d results, got %d", len(tasks), len(results))
	}

	// Проверяем, что все задачи завершились успешно
	for i, result := range results {
		if !strings.Contains(result, "completed") {
			t.Errorf("Task %d should be completed, got: %s", i+1, result)
		}
	}
}

func TestProcessTasksWithContextTimeout(t *testing.T) {
	tasks := []LongRunningTask{
		{ID: 1, Duration: 100},
		{ID: 2, Duration: 200},
		{ID: 3, Duration: 300},
	}

	timeout := int64(150)
	results := ProcessTasksWithContext(tasks, timeout)

	if len(results) != len(tasks) {
		t.Errorf("Expected %d results, got %d", len(tasks), len(results))
	}

	// Проверяем, что некоторые задачи были отменены
	cancelledCount := 0
	for _, result := range results {
		if strings.Contains(result, "cancelled") {
			cancelledCount++
		}
	}

	if cancelledCount == 0 {
		t.Error("Expected some tasks to be cancelled due to timeout")
	}
}

func TestProcessTasksWithContextShortTimeout(t *testing.T) {
	tasks := []LongRunningTask{
		{ID: 1, Duration: 100},
		{ID: 2, Duration: 100},
	}

	timeout := int64(10)
	results := ProcessTasksWithContext(tasks, timeout)

	if len(results) != len(tasks) {
		t.Errorf("Expected %d results, got %d", len(tasks), len(results))
	}

	// Все задачи должны быть отменены
	for i, result := range results {
		if !strings.Contains(result, "cancelled") {
			t.Errorf("Task %d should be cancelled, got: %s", i+1, result)
		}
	}
}

func TestProcessTasksWithContextEmptyList(t *testing.T) {
	tasks := []LongRunningTask{}
	timeout := int64(100)

	results := ProcessTasksWithContext(tasks, timeout)

	if len(results) != 0 {
		t.Errorf("Expected 0 results for empty task list, got %d", len(results))
	}
}

func TestProcessTasksWithContextSingleTask(t *testing.T) {
	tasks := []LongRunningTask{
		{ID: 1, Duration: 50},
	}

	timeout := int64(100)
	results := ProcessTasksWithContext(tasks, timeout)

	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}

	if !strings.Contains(results[0], "completed") {
		t.Errorf("Single task should be completed, got: %s", results[0])
	}
}

func TestProcessTasksWithContextZeroTimeout(t *testing.T) {
	tasks := []LongRunningTask{
		{ID: 1, Duration: 10},
		{ID: 2, Duration: 10},
	}

	timeout := int64(0)
	results := ProcessTasksWithContext(tasks, timeout)

	if len(results) != len(tasks) {
		t.Errorf("Expected %d results, got %d", len(tasks), len(results))
	}

	// Все задачи должны быть отменены немедленно
	for i, result := range results {
		if !strings.Contains(result, "cancelled") {
			t.Errorf("Task %d should be cancelled with zero timeout, got: %s", i+1, result)
		}
	}
}
