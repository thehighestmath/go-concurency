package task8

import (
	"testing"
)

func TestRunPipeline(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Pipeline n=1",
			n:        1,
			expected: []int{}, // 1 -> четное? нет -> пустой результат
		},
		{
			name:     "Pipeline n=2",
			n:        2,
			expected: []int{4}, // 1,2 -> 2 -> 4
		},
		{
			name:     "Pipeline n=4",
			n:        4,
			expected: []int{4, 8}, // 1,2,3,4 -> 2,4 -> 4,8
		},
		{
			name:     "Pipeline n=6",
			n:        6,
			expected: []int{4, 8, 12}, // 1,2,3,4,5,6 -> 2,4,6 -> 4,8,12
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RunPipeline(tt.n)
			if !slicesEqual(result, tt.expected) {
				t.Errorf("RunPipeline(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestRunPipelineZero(t *testing.T) {
	result := RunPipeline(0)
	if len(result) != 0 {
		t.Errorf("RunPipeline(0) should return empty slice, got %v", result)
	}
}

func TestRunPipelineNegative(t *testing.T) {
	result := RunPipeline(-5)
	if len(result) != 0 {
		t.Errorf("RunPipeline(-5) should return empty slice, got %v", result)
	}
}

func TestRunPipelineLarge(t *testing.T) {
	n := 20
	result := RunPipeline(n)

	// Проверяем, что результат содержит только четные числа, умноженные на 2
	expectedCount := n / 2 // количество четных чисел от 1 до n
	if len(result) != expectedCount {
		t.Errorf("Expected %d results, got %d", expectedCount, len(result))
	}

	// Проверяем, что все числа четные и больше 0
	for i, num := range result {
		if num <= 0 || num%2 != 0 {
			t.Errorf("Result[%d] = %d should be positive even number", i, num)
		}
	}
}

func TestCreateNumberGenerator(t *testing.T) {
	ch := CreateNumberGenerator(5)

	expected := []int{1, 2, 3, 4, 5}
	var result []int

	for num := range ch {
		result = append(result, num)
	}

	if !slicesEqual(result, expected) {
		t.Errorf("CreateNumberGenerator(5) = %v, expected %v", result, expected)
	}
}

func TestFilterEvenNumbers(t *testing.T) {
	input := make(chan int, 5)
	go func() {
		input <- 1
		input <- 2
		input <- 3
		input <- 4
		input <- 5
		close(input)
	}()

	output := FilterEvenNumbers(input)

	expected := []int{2, 4}
	var result []int

	for num := range output {
		result = append(result, num)
	}

	if !slicesEqual(result, expected) {
		t.Errorf("FilterEvenNumbers = %v, expected %v", result, expected)
	}
}

func TestDoubleNumbers(t *testing.T) {
	input := make(chan int, 3)
	go func() {
		input <- 1
		input <- 2
		input <- 3
		close(input)
	}()

	output := DoubleNumbers(input)

	expected := []int{2, 4, 6}
	var result []int

	for num := range output {
		result = append(result, num)
	}

	if !slicesEqual(result, expected) {
		t.Errorf("DoubleNumbers = %v, expected %v", result, expected)
	}
}

// Вспомогательная функция для сравнения слайсов
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

