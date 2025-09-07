package task2

import (
	"testing"
)

func TestSumNumbersWithChannels(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"Sum 1 to 1", 1, 1},
		{"Sum 1 to 5", 5, 15},       // 1+2+3+4+5 = 15
		{"Sum 1 to 10", 10, 55},     // 1+2+...+10 = 55
		{"Sum 1 to 100", 100, 5050}, // 1+2+...+100 = 5050
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumNumbersWithChannels(tt.n)
			if result != tt.expected {
				t.Errorf("SumNumbersWithChannels(%d) = %d, expected %d", tt.n, result, tt.expected)
			}
		})
	}
}

func TestSumNumbersWithChannelsZero(t *testing.T) {
	result := SumNumbersWithChannels(0)
	if result != 0 {
		t.Errorf("SumNumbersWithChannels(0) = %d, expected 0", result)
	}
}

func TestSumNumbersWithChannelsNegative(t *testing.T) {
	result := SumNumbersWithChannels(-5)
	if result != 0 {
		t.Errorf("SumNumbersWithChannels(-5) = %d, expected 0", result)
	}
}

func TestSumNumbersWithChannelsLarge(t *testing.T) {
	result := SumNumbersWithChannels(1000)
	expected := 1000 * 1001 / 2 // Формула суммы арифметической прогрессии
	if result != expected {
		t.Errorf("SumNumbersWithChannels(1000) = %d, expected %d", result, expected)
	}
}
