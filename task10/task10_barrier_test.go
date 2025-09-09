package task10

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBarrierBasic(t *testing.T) {
	barrier := NewBarrier(3)

	var wg sync.WaitGroup
	results := make([]int, 3)

	// Запускаем 3 горутины
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Работа до barrier
			results[id] = id * 10

			// Ждем на barrier
			barrier.Wait()

			// Работа после barrier
			results[id] += id
		}(i)
	}

	wg.Wait()

	// Проверяем результаты
	expected := []int{0, 11, 22} // (0*10+0), (1*10+1), (2*10+2)
	for i, result := range results {
		if result != expected[i] {
			t.Errorf("Goroutine %d result = %d, expected %d", i, result, expected[i])
		}
	}
}

func TestBarrierSingleGoroutine(t *testing.T) {
	barrier := NewBarrier(1)

	done := make(chan bool)
	go func() {
		barrier.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Ожидаемое поведение
	case <-time.After(1 * time.Second):
		t.Error("Single goroutine barrier should complete immediately")
	}
}

func TestBarrierTwoGoroutines(t *testing.T) {
	barrier := NewBarrier(2)

	var wg sync.WaitGroup
	order := make([]int, 0, 2)
	var mu sync.Mutex

	// Первая горутина
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		order = append(order, 1)
		mu.Unlock()

		barrier.Wait()

		mu.Lock()
		order = append(order, 3)
		mu.Unlock()
	}()

	// Небольшая задержка
	time.Sleep(10 * time.Millisecond)

	// Вторая горутина
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		order = append(order, 2)
		mu.Unlock()

		barrier.Wait()

		mu.Lock()
		order = append(order, 4)
		mu.Unlock()
	}()

	wg.Wait()

	// Проверяем порядок: сначала обе горутины до barrier, потом обе после
	if len(order) != 4 {
		t.Errorf("Expected 4 events, got %d", len(order))
	}

	if order[0] != 1 || order[1] != 2 {
		t.Errorf("Expected order [1,2,...], got %v", order)
	}

	if !(order[2] == 3 && order[3] == 4) && !(order[2] == 4 && order[3] == 3) {
		t.Errorf("Expected order [...,3,4] or [...,4,3], got %v", order)
	}
}

func TestBarrierManyGoroutines(t *testing.T) {
	numGoroutines := 10
	barrier := NewBarrier(numGoroutines)

	var wg sync.WaitGroup
	results := make([]int, numGoroutines)

	start := time.Now()

	// Запускаем много горутин
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Симулируем работу
			time.Sleep(time.Duration(id) * time.Millisecond)
			results[id] = id

			barrier.Wait()

			// Все горутины должны дойти до этой точки одновременно
			results[id] += 1000
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	// Проверяем, что все горутины завершились
	for i, result := range results {
		expected := i + 1000
		if result != expected {
			t.Errorf("Goroutine %d result = %d, expected %d", i, result, expected)
		}
	}

	// Проверяем, что все горутины ждали самую медленную
	// Самая медленная горутина должна была работать ~9ms
	if elapsed < 9*time.Millisecond {
		t.Errorf("Barrier should wait for slowest goroutine, elapsed: %v", elapsed)
	}
}

func TestBarrierZeroGoroutines(t *testing.T) {
	barrier := NewBarrier(0)

	// Barrier с 0 горутин должен работать без проблем
	done := make(chan bool)
	go func() {
		barrier.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Ожидаемое поведение
	case <-time.After(1 * time.Second):
		t.Error("Zero goroutines barrier should complete immediately")
	}
}

func TestBarrierReuse(t *testing.T) {
	barrier := NewBarrier(5)
	const rounds = 10
	cnt := 0
	var mu sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for r := 0; r < rounds; r++ {
				barrier.Wait()
				mu.Lock()
				cnt++
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	if cnt != 5*rounds {
		t.Errorf("Expected cnt = %d, got %d", 5*rounds, cnt)
	}
}

func TestBarrierSequentialPhases(t *testing.T) {
	n := 4
	barrier := NewBarrier(n)
	phases := 50
	var mu sync.Mutex
	counts := make([]int, phases)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for phase := 0; phase < phases; phase++ {
				barrier.Wait()
				if phase == 0 {
					// First phase, nothing to check
					continue
				}
				mu.Lock()
				counts[phase]++
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	for i, v := range counts[1:] { // phase 0 не трогаем
		if v != n {
			t.Errorf("Phase %d: count expected %d, got %d", i+1, n, v)
		}
	}
}

func TestBarrierRacyPhases(t *testing.T) {
	n := 3
	barrier := NewBarrier(n)
	done := make(chan bool, 2)
	// Первая "волна"
	for i := 0; i < n; i++ {
		go func() {
			barrier.Wait()
			done <- true
		}()
	}
	// Вторая "волна" — очень быстро за ней
	for i := 0; i < n; i++ {
		go func() {
			barrier.Wait()
			done <- true
		}()
	}

	timeout := time.After(1 * time.Second)
	for i := 0; i < 2*n; i++ {
		select {
		case <-done:
			// ok
		case <-timeout:
			t.Fatalf("Barrier deadlock or phase confusion detected")
		}
	}
}

func TestBarrierNegative(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Should panic for negative N")
		}
	}()
	NewBarrier(-1)
}

func TestBarrierNoRaceFunny(t *testing.T) {
	barrier := NewBarrier(2)
	for i := 0; i < 100; i++ {
		done := make(chan bool, 2)
		go func() {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
			barrier.Wait()
			done <- true
		}()
		go func() {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
			barrier.Wait()
			done <- true
		}()
		<-done
		<-done
	}
}
