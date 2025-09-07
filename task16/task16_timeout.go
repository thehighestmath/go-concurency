package task16

// Task 16: Timeout Pattern
// Реализуйте различные паттерны работы с таймаутами.

// "context"
// "fmt"
// "sync"
// "time"

// TimeoutManager управляет таймаутами
type TimeoutManager struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map активных таймаутов
	// - mutex для синхронизации
	// - канал для уведомлений
}

// TimeoutConfig содержит настройки таймаута
type TimeoutConfig struct {
	Duration int64  // Длительность в миллисекундах
	Callback func() // Функция обратного вызова
	ID       string // Уникальный идентификатор
}

// NewTimeoutManager создает новый менеджер таймаутов
func NewTimeoutManager() *TimeoutManager {
	// TODO: Реализуйте конструктор
	return nil
}

// SetTimeout устанавливает таймаут
func (tm *TimeoutManager) SetTimeout(config TimeoutConfig) {
	// TODO: Реализуйте метод
	// 1. Создайте горутину с таймером
	// 2. Сохраните таймаут в map
	// 3. По истечении времени вызовите callback
}

// ClearTimeout отменяет таймаут
func (tm *TimeoutManager) ClearTimeout(id string) {
	// TODO: Реализуйте метод
	// Удалите таймаут из map
}

// ClearAllTimeouts отменяет все таймауты
func (tm *TimeoutManager) ClearAllTimeouts() {
	// TODO: Реализуйте метод
}

// GetActiveTimeouts возвращает количество активных таймаутов
func (tm *TimeoutManager) GetActiveTimeouts() int {
	// TODO: Реализуйте метод
	return 0
}

// ExecuteWithTimeout выполняет функцию с таймаутом
func ExecuteWithTimeout(fn func() error, timeout int64) error {
	// TODO: Реализуйте функцию
	// 1. Создайте context с таймаутом
	// 2. Запустите функцию в горутине
	// 3. Ждите завершения или таймаута
	// 4. Верните результат или ошибку таймаута
	return nil
}

// ExecuteWithDeadline выполняет функцию с дедлайном
func ExecuteWithDeadline(fn func() error, deadline int64) error {
	// TODO: Реализуйте функцию
	// Аналогично ExecuteWithTimeout, но с абсолютным временем
	return nil
}

