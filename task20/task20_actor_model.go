package task20

// Task 20: Actor Model
// Реализуйте модель Actor для изоляции состояния и обмена сообщениями.

// "fmt"
// "sync"

// Message представляет сообщение для Actor
type Message struct {
	Type    string
	Data    interface{}
	ReplyTo chan interface{}
}

// Actor представляет актора в модели Actor
type Actor struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - канал для получения сообщений
	// - состояние актора
	// - функция обработки сообщений
	// - канал для остановки
}

// ActorSystem управляет акторами
type ActorSystem struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map акторов по ID
	// - mutex для синхронизации
	// - канал для остановки системы
}

// NewActorSystem создает новую систему акторов
func NewActorSystem() *ActorSystem {
	// TODO: Реализуйте конструктор
	return nil
}

// SpawnActor создает нового актора
func (as *ActorSystem) SpawnActor(id string, initialState interface{}, handler func(interface{}, Message) interface{}) *Actor {
	// TODO: Реализуйте метод
	// 1. Создайте канал для сообщений
	// 2. Создайте актора с обработчиком
	// 3. Запустите горутину для обработки сообщений
	// 4. Добавьте в систему
	return nil
}

// SendMessage отправляет сообщение актору
func (as *ActorSystem) SendMessage(actorID string, msg Message) {
	// TODO: Реализуйте метод
	// 1. Найдите актора по ID
	// 2. Отправьте сообщение в его канал
}

// SendMessageAsync отправляет сообщение асинхронно
func (as *ActorSystem) SendMessageAsync(actorID string, msgType string, data interface{}) {
	// TODO: Реализуйте метод
}

// SendMessageSync отправляет сообщение синхронно
func (as *ActorSystem) SendMessageSync(actorID string, msgType string, data interface{}) interface{} {
	// TODO: Реализуйте метод
	// 1. Создайте канал для ответа
	// 2. Отправьте сообщение
	// 3. Ждите ответ
	return nil
}

// StopActor останавливает актора
func (as *ActorSystem) StopActor(actorID string) {
	// TODO: Реализуйте метод
}

// StopAllActors останавливает всех акторов
func (as *ActorSystem) StopAllActors() {
	// TODO: Реализуйте метод
}

// GetActorCount возвращает количество акторов
func (as *ActorSystem) GetActorCount() int {
	// TODO: Реализуйте метод
	return 0
}

// CounterActor представляет актора-счетчик
type CounterActor struct {
	*Actor
	Count int
}

// NewCounterActor создает нового актора-счетчик
func NewCounterActor(system *ActorSystem, id string) *CounterActor {
	// TODO: Реализуйте конструктор
	// 1. Создайте обработчик сообщений
	// 2. Запустите актора через систему
	return nil
}

// Increment увеличивает счетчик
func (ca *CounterActor) Increment() {
	// TODO: Реализуйте метод
}

// Decrement уменьшает счетчик
func (ca *CounterActor) Decrement() {
	// TODO: Реализуйте метод
}

// GetValue возвращает значение счетчика
func (ca *CounterActor) GetValue() int {
	// TODO: Реализуйте метод
	return 0
}

// Reset сбрасывает счетчик
func (ca *CounterActor) Reset() {
	// TODO: Реализуйте метод
}

