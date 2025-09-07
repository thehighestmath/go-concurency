package task4

import (
	"testing"
	// "time"
)

func TestSumFromTwoChannels(t *testing.T) {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	// Отправляем числа в каналы
	go func() {
		ch1 <- 1
		ch1 <- 3
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		ch2 <- 2
		ch2 <- 4
		ch2 <- 6
		close(ch2)
	}()

	result := SumFromTwoChannels(ch1, ch2, 1000)
	expected := 1 + 3 + 5 + 2 + 4 + 6 // 21

	if result != expected {
		t.Errorf("Expected sum %d, got %d", expected, result)
	}
}

func TestSumFromTwoChannelsTimeout(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Не отправляем ничего в каналы, чтобы сработал timeout
	result := SumFromTwoChannels(ch1, ch2, 100)

	if result != 0 {
		t.Errorf("Expected 0 on timeout, got %d", result)
	}
}

func TestSumFromTwoChannelsOneEmpty(t *testing.T) {
	ch1 := make(chan int, 3)
	ch2 := make(chan int)

	// Отправляем только в первый канал
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()

	result := SumFromTwoChannels(ch1, ch2, 100)
	expected := 1 + 2 + 3 // 6

	if result != expected {
		t.Errorf("Expected sum %d, got %d", expected, result)
	}
}

func TestSumFromTwoChannelsBothEmpty(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Закрываем каналы сразу
	close(ch1)
	close(ch2)

	result := SumFromTwoChannels(ch1, ch2, 100)

	if result != 0 {
		t.Errorf("Expected 0 for empty channels, got %d", result)
	}
}

func TestSumFromTwoChannelsLargeNumbers(t *testing.T) {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	// Отправляем большие числа
	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i * 1000
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			ch2 <- i * 2000
		}
		close(ch2)
	}()

	result := SumFromTwoChannels(ch1, ch2, 1000)
	expected := (1000 + 2000 + 3000 + 4000 + 5000) + (2000 + 4000 + 6000 + 8000 + 10000) // 45000

	if result != expected {
		t.Errorf("Expected sum %d, got %d", expected, result)
	}
}
