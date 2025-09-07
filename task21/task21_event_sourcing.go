package task21

// Task 21: Event Sourcing Pattern
// Реализуйте паттерн Event Sourcing для хранения состояния как последовательности событий.

// "fmt"
// "sync"
// "time"

// Event представляет событие в системе
type Event struct {
	ID        string
	Type      string
	Data      interface{}
	Timestamp int64
	Version   int
}

// EventStore хранит события
type EventStore struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map событий по агрегату
	// - mutex для синхронизации
	// - счетчик версий
}

// Aggregate представляет агрегат в DDD
type Aggregate struct {
	ID      string
	Version int
	State   interface{}
}

// NewEventStore создает новое хранилище событий
func NewEventStore() *EventStore {
	// TODO: Реализуйте конструктор
	return nil
}

// AppendEvents добавляет события в хранилище
func (es *EventStore) AppendEvents(aggregateID string, events []Event) error {
	// TODO: Реализуйте метод
	// 1. Проверьте версию агрегата
	// 2. Добавьте события
	// 3. Обновите версию
	return nil
}

// GetEvents возвращает события для агрегата
func (es *EventStore) GetEvents(aggregateID string) ([]Event, error) {
	// TODO: Реализуйте метод
	return nil, nil
}

// GetEventsFromVersion возвращает события с определенной версии
func (es *EventStore) GetEventsFromVersion(aggregateID string, fromVersion int) ([]Event, error) {
	// TODO: Реализуйте метод
	return nil, nil
}

// GetAggregateVersion возвращает версию агрегата
func (es *EventStore) GetAggregateVersion(aggregateID string) int {
	// TODO: Реализуйте метод
	return 0
}

// EventHandler обрабатывает события
type EventHandler struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map обработчиков по типу события
	// - mutex для синхронизации
}

// NewEventHandler создает новый обработчик событий
func NewEventHandler() *EventHandler {
	// TODO: Реализуйте конструктор
	return nil
}

// RegisterHandler регистрирует обработчик для типа события
func (eh *EventHandler) RegisterHandler(eventType string, handler func(Event)) {
	// TODO: Реализуйте метод
}

// HandleEvent обрабатывает событие
func (eh *EventHandler) HandleEvent(event Event) {
	// TODO: Реализуйте метод
	// 1. Найдите обработчик по типу события
	// 2. Вызовите обработчик
}

// HandleEvents обрабатывает массив событий
func (eh *EventHandler) HandleEvents(events []Event) {
	// TODO: Реализуйте метод
}

// EventBus управляет событиями
type EventBus struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - канал для событий
	// - обработчики событий
	// - горутины для обработки
}

// NewEventBus создает новую шину событий
func NewEventBus() *EventBus {
	// TODO: Реализуйте конструктор
	return nil
}

// PublishEvent публикует событие
func (eb *EventBus) PublishEvent(event Event) {
	// TODO: Реализуйте метод
}

// Subscribe подписывается на события
func (eb *EventBus) Subscribe(eventType string, handler func(Event)) {
	// TODO: Реализуйте метод
}

// Start запускает шину событий
func (eb *EventBus) Start() {
	// TODO: Реализуйте метод
	// Запустите горутины для обработки событий
}

// Stop останавливает шину событий
func (eb *EventBus) Stop() {
	// TODO: Реализуйте метод
}

// BankAccountAggregate представляет агрегат банковского счета
type BankAccountAggregate struct {
	ID      string
	Balance int
	Version int
}

// NewBankAccountAggregate создает новый агрегат счета
func NewBankAccountAggregate(id string) *BankAccountAggregate {
	// TODO: Реализуйте конструктор
	return nil
}

// ApplyEvent применяет событие к агрегату
func (ba *BankAccountAggregate) ApplyEvent(event Event) {
	// TODO: Реализуйте метод
	// 1. Проверьте тип события
	// 2. Обновите состояние
	// 3. Увеличьте версию
}

// Deposit создает событие депозита
func (ba *BankAccountAggregate) Deposit(amount int) Event {
	// TODO: Реализуйте метод
	return Event{}
}

// Withdraw создает событие вывода
func (ba *BankAccountAggregate) Withdraw(amount int) Event {
	// TODO: Реализуйте метод
	return Event{}
}

