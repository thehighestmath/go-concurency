package task10

import (
	"math/rand"
	"sync"
	"time"
)

// Task 10: Barrier Synchronization
// Реализуйте barrier - механизм синхронизации, который заставляет
// все горутины ждать друг друга перед продолжением выполнения.

// Barrier обеспечивает синхронизацию горутин
type Barrier struct {
	mu         *sync.Mutex
	cond       *sync.Cond
	waitCount  int
	totalCount int
}

// NewBarrier создает новый barrier для n горутин
func NewBarrier(n int) *Barrier {
	if n < 0 {
		panic(n)
	}
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	return &Barrier{
		mu:         mu,
		cond:       cond,
		waitCount:  0,
		totalCount: n,
	}
}

// Wait блокирует горутину до тех пор, пока все n горутин не вызовут Wait
func (b *Barrier) Wait() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.totalCount > 0 {
		b.waitCount++
	}
	if b.totalCount != b.waitCount {
		b.cond.Wait()
	} else {
		b.cond.Broadcast()
		b.waitCount = 0
	}
	// TODO: Реализуйте метод
	// 1. Заблокируйте mutex
	// 2. Увеличьте счетчик ожидающих
	// 3. Если не все горутины пришли - ждите на cond
	// 4. Если все пришли - разбудите всех и сбросьте счетчик
	// 5. Разблокируйте mutex
}

// BarrierExample демонстрирует использование barrier
func BarrierExample(numGoroutines int) []int {
	b := NewBarrier(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			time.Sleep(time.Duration(100*rand.Float64()) * time.Millisecond)
			b.Wait()
			time.Sleep(time.Duration(100*rand.Float64()) * time.Millisecond)
		}()
	}

	out := make([]int, numGoroutines)
	for i := range out {
		out[i] = 1
	}

	return out
}
