package task15

// Task 15: Retry Pattern with Exponential Backoff
// Реализуйте паттерн Retry с экспоненциальной задержкой и jitter.

// "context"
// "fmt"
// "math/rand"
// "sync"
// "time"

// RetryConfig содержит настройки retry
type RetryConfig struct {
	MaxAttempts   int     // Максимальное количество попыток
	BaseDelay     int64   // Базовая задержка в миллисекундах
	MaxDelay      int64   // Максимальная задержка в миллисекундах
	BackoffFactor float64 // Коэффициент экспоненциального роста
	Jitter        bool    // Добавлять ли случайность к задержке
}

// RetryableFunc представляет функцию, которую можно повторить
type RetryableFunc func() error

// RetryManager управляет повторными попытками
type RetryManager struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - конфигурация
	// - статистика попыток
	// - mutex для синхронизации
}

// NewRetryManager создает новый менеджер retry
func NewRetryManager(config RetryConfig) *RetryManager {
	// TODO: Реализуйте конструктор
	return nil
}

// Execute выполняет функцию с повторными попытками
func (rm *RetryManager) Execute(ctx interface{}, fn RetryableFunc) error {
	// TODO: Реализуйте метод
	// 1. Выполните функцию
	// 2. Если ошибка - вычислите задержку
	// 3. Добавьте jitter если нужно
	// 4. Подождите задержку
	// 5. Повторите до MaxAttempts
	return nil
}

// ExecuteWithCustomRetry выполняет функцию с кастомной логикой retry
func (rm *RetryManager) ExecuteWithCustomRetry(ctx interface{}, fn RetryableFunc, shouldRetry func(error) bool) error {
	// TODO: Реализуйте метод
	// Аналогично Execute, но с кастомной функцией shouldRetry
	return nil
}

// GetStats возвращает статистику retry
func (rm *RetryManager) GetStats() map[string]int {
	// TODO: Реализуйте метод
	// Верните map с ключами: "attempts", "successes", "failures"
	return nil
}

// ResetStats сбрасывает статистику
func (rm *RetryManager) ResetStats() {
	// TODO: Реализуйте метод
}

