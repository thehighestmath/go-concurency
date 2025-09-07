package task11

// Task 11: Semaphore
// Реализуйте семафор - механизм для ограничения количества
// одновременно выполняющихся операций.

// "context"
// "fmt"
// "sync"

// Semaphore ограничивает количество одновременных операций
type Semaphore struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобится канал для токенов
}

// NewSemaphore создает новый семафор с указанной емкостью
func NewSemaphore(capacity int) *Semaphore {
	// TODO: Реализуйте конструктор
	// Создайте канал с буфером capacity и заполните его токенами
	return nil
}

// Acquire получает токен из семафора
func (s *Semaphore) Acquire() {
	// TODO: Реализуйте метод
	// Получите токен из канала (блокирующая операция)
}

// TryAcquire пытается получить токен без блокировки
func (s *Semaphore) TryAcquire() bool {
	// TODO: Реализуйте метод
	// Попробуйте получить токен из канала (неблокирующая операция)
	// Верните true если получили, false если нет
	return false
}

// AcquireWithContext получает токен с учетом контекста
func (s *Semaphore) AcquireWithContext(ctx interface{}) error {
	// TODO: Реализуйте метод
	// Ждите токен с учетом контекста
	// Верните ошибку если контекст отменен
	return nil
}

// Release возвращает токен в семафор
func (s *Semaphore) Release() {
	// TODO: Реализуйте метод
	// Верните токен в канал
}

// SemaphoreExample демонстрирует использование семафора
func SemaphoreExample(numWorkers, maxConcurrent int) []string {
	// TODO: Реализуйте эту функцию
	// 1. Создайте семафор с емкостью maxConcurrent
	// 2. Запустите numWorkers горутин
	// 3. Каждая горутина должна:
	//    - получить токен из семафора
	//    - выполнить работу (симулировать задержку)
	//    - вернуть токен в семафор
	// 4. Соберите результаты от всех горутин
	// 5. Верните слайс результатов

	return nil
}
