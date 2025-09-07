package task6

import (
	"testing"
	"time"
)

func TestFanOutFanIn(t *testing.T) {
	tests := []struct {
		name        string
		numbers     []int
		numWorkers  int
		processFunc func(int) int
		expected    []int
	}{
		{
			name:        "Single worker, square function",
			numbers:     []int{1, 2, 3, 4, 5},
			numWorkers:  1,
			processFunc: func(x int) int { return x * x },
			expected:    []int{1, 4, 9, 16, 25},
		},
		{
			name:        "Multiple workers, double function",
			numbers:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			numWorkers:  3,
			processFunc: func(x int) int { return x * 2 },
			expected:    []int{2, 4, 6, 8, 10, 12, 14, 16},
		},
		{
			name:        "Two workers, increment function",
			numbers:     []int{10, 20, 30, 40},
			numWorkers:  2,
			processFunc: func(x int) int { return x + 1 },
			expected:    []int{11, 21, 31, 41},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем входной канал
			input := make(chan int)
			go func() {
				for _, num := range tt.numbers {
					input <- num
				}
				close(input)
			}()

			// Запускаем FanOutFanIn
			output := FanOutFanIn(input, tt.numWorkers, tt.processFunc)

			// Собираем результаты
			var results []int
			for result := range output {
				results = append(results, result)
			}

			// Проверяем количество результатов
			if len(results) != len(tt.expected) {
				t.Errorf("Expected %d results, got %d", len(tt.expected), len(results))
				return
			}

			// Проверяем, что все ожидаемые значения присутствуют
			// (порядок может отличаться из-за параллельности)
			expectedMap := make(map[int]int)
			for _, val := range tt.expected {
				expectedMap[val]++
			}

			resultMap := make(map[int]int)
			for _, val := range results {
				resultMap[val]++
			}

			for val, count := range expectedMap {
				if resultMap[val] != count {
					t.Errorf("Expected %d occurrences of %d, got %d", count, val, resultMap[val])
				}
			}
		})
	}
}

func TestFanOutFanInEmptyInput(t *testing.T) {
	input := make(chan int)
	close(input)

	output := FanOutFanIn(input, 2, func(x int) int { return x * 2 })

	// Должен быть пустой результат
	var results []int
	for result := range output {
		results = append(results, result)
	}

	if len(results) != 0 {
		t.Errorf("Expected 0 results for empty input, got %d", len(results))
	}
}

func TestFanOutFanInZeroWorkers(t *testing.T) {
	input := make(chan int)
	go func() {
		input <- 1
		input <- 2
		close(input)
	}()

	// Тестируем с 0 воркерами - должно работать без паники
	output := FanOutFanIn(input, 0, func(x int) int { return x * 2 })

	var results []int
	for result := range output {
		results = append(results, result)
	}

	// С 0 воркерами не должно быть результатов
	if len(results) != 0 {
		t.Errorf("Expected 0 results with 0 workers, got %d", len(results))
	}
}

func TestFanOutFanInConcurrent(t *testing.T) {
	// Тест на конкурентность - проверяем, что все воркеры работают
	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = i + 1
	}

	input := make(chan int)
	go func() {
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	output := FanOutFanIn(input, 5, func(x int) int { return x * x })

	var results []int
	for result := range output {
		results = append(results, result)
	}

	if len(results) != 100 {
		t.Errorf("Expected 100 results, got %d", len(results))
	}

	// Проверяем, что все квадраты чисел от 1 до 100 присутствуют
	expectedSquares := make(map[int]bool)
	for i := 1; i <= 100; i++ {
		expectedSquares[i*i] = true
	}

	for _, result := range results {
		if !expectedSquares[result] {
			t.Errorf("Unexpected result: %d", result)
		}
	}
}

func TestFanOutFanInWithDelay(t *testing.T) {
	// Тест с задержкой в функции обработки
	input := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	start := time.Now()
	output := FanOutFanIn(input, 2, func(x int) int {
		time.Sleep(10 * time.Millisecond)
		return x * 2
	})

	var results []int
	for result := range output {
		results = append(results, result)
	}
	elapsed := time.Since(start)

	// С 2 воркерами и 5 элементами, время должно быть примерно 30ms
	// (3 батча по 10ms каждый)
	if elapsed < 20*time.Millisecond || elapsed > 50*time.Millisecond {
		t.Errorf("Expected execution time around 30ms, got %v", elapsed)
	}

	if len(results) != 5 {
		t.Errorf("Expected 5 results, got %d", len(results))
	}
}
