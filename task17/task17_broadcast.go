package task17

// Task 17: Broadcast Pattern
// Реализуйте паттерн Broadcast для уведомления множества получателей.

// "fmt"
// "sync"

// BroadcastChannel представляет канал для broadcast
type BroadcastChannel struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map подписчиков
	// - mutex для синхронизации
	// - канал для закрытия
}

// Subscriber представляет подписчика
type Subscriber struct {
	ID   string
	Ch   chan interface{}
	Done chan struct{}
}

// NewBroadcastChannel создает новый broadcast канал
func NewBroadcastChannel() *BroadcastChannel {
	// TODO: Реализуйте конструктор
	return nil
}

// Subscribe подписывается на broadcast
func (bc *BroadcastChannel) Subscribe(id string) *Subscriber {
	// TODO: Реализуйте метод
	// 1. Создайте нового подписчика
	// 2. Добавьте его в map подписчиков
	// 3. Верните подписчика
	return nil
}

// Unsubscribe отписывается от broadcast
func (bc *BroadcastChannel) Unsubscribe(id string) {
	// TODO: Реализуйте метод
	// 1. Найдите подписчика по ID
	// 2. Закройте его канал
	// 3. Удалите из map
}

// Broadcast отправляет сообщение всем подписчикам
func (bc *BroadcastChannel) Broadcast(message interface{}) {
	// TODO: Реализуйте метод
	// 1. Пройдитесь по всем подписчикам
	// 2. Отправьте сообщение в каждый канал
	// 3. Обработайте заблокированные каналы
}

// GetSubscriberCount возвращает количество подписчиков
func (bc *BroadcastChannel) GetSubscriberCount() int {
	// TODO: Реализуйте метод
	return 0
}

// Close закрывает broadcast канал
func (bc *BroadcastChannel) Close() {
	// TODO: Реализуйте метод
	// 1. Отпишите всех подписчиков
	// 2. Закройте канал
}

// BroadcastManager управляет несколькими broadcast каналами
type BroadcastManager struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map каналов по темам
	// - mutex для синхронизации
}

// NewBroadcastManager создает новый менеджер broadcast
func NewBroadcastManager() *BroadcastManager {
	// TODO: Реализуйте конструктор
	return nil
}

// GetChannel возвращает канал по теме
func (bm *BroadcastManager) GetChannel(topic string) *BroadcastChannel {
	// TODO: Реализуйте метод
	// Создайте канал если его нет
	return nil
}

// BroadcastToTopic отправляет сообщение в канал по теме
func (bm *BroadcastManager) BroadcastToTopic(topic string, message interface{}) {
	// TODO: Реализуйте метод
}

