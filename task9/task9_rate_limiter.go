package task9

// Task 9: Rate Limiter
// Реализуйте rate limiter с использованием token bucket алгоритма.
// Rate limiter должен ограничивать количество запросов в единицу времени.

// "context"
// "fmt"
// "sync"
// "time"

// RateLimiter ограничивает количество запросов
type RateLimiter struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - канал для токенов
	// - ticker для пополнения токенов
	// - mutex для синхронизации
	// - контекст для остановки
}

// NewRateLimiter создает новый rate limiter
// rate - количество запросов в секунду
// burst - максимальное количество токенов
func NewRateLimiter(rate int, burst int) *RateLimiter {
	// TODO: Реализуйте конструктор
	// 1. Создайте канал с буфером burst
	// 2. Заполните канал токенами
	// 3. Запустите горутину для пополнения токенов
	return nil
}

// Allow проверяет, можно ли выполнить запрос
func (rl *RateLimiter) Allow() bool {
	// TODO: Реализуйте метод
	// Попробуйте получить токен из канала
	// Верните true если токен получен, false если нет
	return false
}

// Wait ждет, пока не станет доступен токен
func (rl *RateLimiter) Wait(ctx interface{}) error {
	// TODO: Реализуйте метод
	// Ждите токен с учетом контекста
	// Верните ошибку если контекст отменен
	return nil
}

// Stop останавливает rate limiter
func (rl *RateLimiter) Stop() {
	// TODO: Реализуйте метод
	// Остановите ticker и закройте канал
}
