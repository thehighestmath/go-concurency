package task11

import "context"

// Task 11: Semaphore
// Реализуйте семафор - механизм для ограничения количества
// одновременно выполняющихся операций.

// Semaphore ограничивает количество одновременных операций
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore создает новый семафор с указанной емкостью
func NewSemaphore(capacity int) *Semaphore {
	ch := make(chan struct{}, capacity)
	for i := 0; i < capacity; i++ {
		ch <- struct{}{}
	}
	return &Semaphore{
		ch: ch,
	}
}

// Acquire получает токен из семафора
func (s *Semaphore) Acquire() {
	<-s.ch
}

// TryAcquire пытается получить токен без блокировки
func (s *Semaphore) TryAcquire() bool {
	select {
	case <-s.ch:
		return true
	default:
		return false
	}
}

// AcquireWithContext получает токен с учетом контекста
func (s *Semaphore) AcquireWithContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.ch:
		return nil
	}
}

// Release возвращает токен в семафор
func (s *Semaphore) Release() {
	s.ch <- struct{}{}
}
