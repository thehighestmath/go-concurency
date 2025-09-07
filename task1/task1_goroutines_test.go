package task1

import (
	"testing"
	"time"
)

func TestRunNGoroutines(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected time.Duration
	}{
		{"Single goroutine", 1, 100 * time.Millisecond},
		{"Multiple goroutines", 5, 100 * time.Millisecond},
		{"Many goroutines", 10, 100 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			RunNGoroutines(tt.n)
			elapsed := time.Since(start)

			// Проверяем, что функция завершилась
			if elapsed > 200*time.Millisecond {
				t.Errorf("Function took too long: %v", elapsed)
			}

			// Проверяем, что функция не завершилась слишком быстро
			if elapsed < 1*time.Millisecond {
				t.Errorf("Function completed too quickly: %v", elapsed)
			}
		})
	}
}

func TestRunNGoroutinesZero(t *testing.T) {
	start := time.Now()
	RunNGoroutines(0)
	elapsed := time.Since(start)

	// Нулевое количество горутин должно завершиться мгновенно
	if elapsed > 10*time.Millisecond {
		t.Errorf("Zero goroutines should complete instantly, took: %v", elapsed)
	}
}

func TestRunNGoroutinesNegative(t *testing.T) {
	start := time.Now()
	RunNGoroutines(-1)
	elapsed := time.Since(start)

	// Отрицательное количество горутин должно завершиться мгновенно
	if elapsed > 10*time.Millisecond {
		t.Errorf("Negative goroutines should complete instantly, took: %v", elapsed)
	}
}
