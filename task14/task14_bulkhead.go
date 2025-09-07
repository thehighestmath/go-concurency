package task14

// Task 14: Bulkhead Pattern
// Реализуйте паттерн Bulkhead для изоляции ресурсов.
// Разные типы операций должны использовать отдельные пулы ресурсов.

// "context"
// "fmt"
// "sync"
// "time"

// OperationType представляет тип операции
type OperationType int

const (
	OperationRead OperationType = iota
	OperationWrite
	OperationDelete
)

// BulkheadPool управляет пулом ресурсов для определенного типа операций
type BulkheadPool struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - семафор для ограничения ресурсов
	// - тип операции
	// - статистика использования
}

// BulkheadManager управляет несколькими пулами ресурсов
type BulkheadManager struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map пулов по типам операций
	// - mutex для синхронизации
}

// NewBulkheadManager создает новый менеджер bulkhead
func NewBulkheadManager() *BulkheadManager {
	// TODO: Реализуйте конструктор
	// Создайте пулы для каждого типа операций
	return nil
}

// ConfigurePool настраивает пул для определенного типа операций
func (bm *BulkheadManager) ConfigurePool(opType OperationType, maxConcurrency int) {
	// TODO: Реализуйте метод
}

// Execute выполняет операцию через соответствующий пул
func (bm *BulkheadManager) Execute(opType OperationType, fn func() error) error {
	// TODO: Реализуйте метод
	// 1. Получите пул для типа операции
	// 2. Получите ресурс из пула
	// 3. Выполните функцию
	// 4. Освободите ресурс
	return nil
}

// GetPoolStats возвращает статистику пула
func (bm *BulkheadManager) GetPoolStats(opType OperationType) map[string]int {
	// TODO: Реализуйте метод
	// Верните map с ключами: "active", "waiting", "total"
	return nil
}

// GetTotalStats возвращает общую статистику всех пулов
func (bm *BulkheadManager) GetTotalStats() map[string]int {
	// TODO: Реализуйте метод
	return nil
}

