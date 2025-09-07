package task21

import (
	"testing"
	"time"
)

func TestEventStoreBasic(t *testing.T) {
	store := NewEventStore()

	// Создаем события
	events := []Event{
		{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1},
		{ID: "2", Type: "withdraw", Data: 50, Timestamp: time.Now().UnixMilli(), Version: 2},
	}

	// Добавляем события
	err := store.AppendEvents("account1", events)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Проверяем версию агрегата
	version := store.GetAggregateVersion("account1")
	if version != 2 {
		t.Errorf("Expected version 2, got %d", version)
	}
}

func TestEventStoreGetEvents(t *testing.T) {
	store := NewEventStore()

	// Создаем события
	events := []Event{
		{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1},
		{ID: "2", Type: "withdraw", Data: 50, Timestamp: time.Now().UnixMilli(), Version: 2},
	}

	// Добавляем события
	store.AppendEvents("account1", events)

	// Получаем события
	retrievedEvents, err := store.GetEvents("account1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(retrievedEvents) != 2 {
		t.Errorf("Expected 2 events, got %d", len(retrievedEvents))
	}

	if retrievedEvents[0].Type != "deposit" {
		t.Errorf("Expected first event type 'deposit', got %s", retrievedEvents[0].Type)
	}

	if retrievedEvents[1].Type != "withdraw" {
		t.Errorf("Expected second event type 'withdraw', got %s", retrievedEvents[1].Type)
	}
}

func TestEventStoreGetEventsFromVersion(t *testing.T) {
	store := NewEventStore()

	// Создаем события
	events := []Event{
		{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1},
		{ID: "2", Type: "withdraw", Data: 50, Timestamp: time.Now().UnixMilli(), Version: 2},
		{ID: "3", Type: "deposit", Data: 25, Timestamp: time.Now().UnixMilli(), Version: 3},
	}

	// Добавляем события
	store.AppendEvents("account1", events)

	// Получаем события с версии 2
	retrievedEvents, err := store.GetEventsFromVersion("account1", 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(retrievedEvents) != 2 {
		t.Errorf("Expected 2 events from version 2, got %d", len(retrievedEvents))
	}

	if retrievedEvents[0].Type != "withdraw" {
		t.Errorf("Expected first event type 'withdraw', got %s", retrievedEvents[0].Type)
	}

	if retrievedEvents[1].Type != "deposit" {
		t.Errorf("Expected second event type 'deposit', got %s", retrievedEvents[1].Type)
	}
}

func TestEventHandlerBasic(t *testing.T) {
	handler := NewEventHandler()

	// Регистрируем обработчик
	handledEvents := make([]Event, 0)
	handler.RegisterHandler("deposit", func(event Event) {
		handledEvents = append(handledEvents, event)
	})

	// Обрабатываем событие
	event := Event{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1}
	handler.HandleEvent(event)

	// Проверяем, что событие обработано
	if len(handledEvents) != 1 {
		t.Errorf("Expected 1 handled event, got %d", len(handledEvents))
	}

	if handledEvents[0].Type != "deposit" {
		t.Errorf("Expected handled event type 'deposit', got %s", handledEvents[0].Type)
	}
}

func TestEventHandlerMultipleEvents(t *testing.T) {
	handler := NewEventHandler()

	// Регистрируем обработчик
	handledEvents := make([]Event, 0)
	handler.RegisterHandler("deposit", func(event Event) {
		handledEvents = append(handledEvents, event)
	})

	// Обрабатываем несколько событий
	events := []Event{
		{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1},
		{ID: "2", Type: "deposit", Data: 50, Timestamp: time.Now().UnixMilli(), Version: 2},
		{ID: "3", Type: "withdraw", Data: 25, Timestamp: time.Now().UnixMilli(), Version: 3},
	}

	handler.HandleEvents(events)

	// Проверяем, что обработаны только события типа "deposit"
	if len(handledEvents) != 2 {
		t.Errorf("Expected 2 handled events, got %d", len(handledEvents))
	}
}

func TestEventBusBasic(t *testing.T) {
	bus := NewEventBus()

	// Регистрируем обработчик
	handledEvents := make([]Event, 0)
	bus.Subscribe("deposit", func(event Event) {
		handledEvents = append(handledEvents, event)
	})

	// Запускаем шину событий
	bus.Start()
	defer bus.Stop()

	// Публикуем событие
	event := Event{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1}
	bus.PublishEvent(event)

	// Ждем обработки
	time.Sleep(10 * time.Millisecond)

	// Проверяем, что событие обработано
	if len(handledEvents) != 1 {
		t.Errorf("Expected 1 handled event, got %d", len(handledEvents))
	}
}

func TestBankAccountAggregateBasic(t *testing.T) {
	account := NewBankAccountAggregate("account1")

	if account.ID != "account1" {
		t.Errorf("Expected ID 'account1', got %s", account.ID)
	}

	if account.Balance != 0 {
		t.Errorf("Expected initial balance 0, got %d", account.Balance)
	}

	if account.Version != 0 {
		t.Errorf("Expected initial version 0, got %d", account.Version)
	}
}

func TestBankAccountAggregateApplyEvent(t *testing.T) {
	account := NewBankAccountAggregate("account1")

	// Применяем событие депозита
	depositEvent := Event{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1}
	account.ApplyEvent(depositEvent)

	if account.Balance != 100 {
		t.Errorf("Expected balance 100, got %d", account.Balance)
	}

	if account.Version != 1 {
		t.Errorf("Expected version 1, got %d", account.Version)
	}

	// Применяем событие вывода
	withdrawEvent := Event{ID: "2", Type: "withdraw", Data: 50, Timestamp: time.Now().UnixMilli(), Version: 2}
	account.ApplyEvent(withdrawEvent)

	if account.Balance != 50 {
		t.Errorf("Expected balance 50, got %d", account.Balance)
	}

	if account.Version != 2 {
		t.Errorf("Expected version 2, got %d", account.Version)
	}
}

func TestBankAccountAggregateCreateEvents(t *testing.T) {
	account := NewBankAccountAggregate("account1")

	// Создаем событие депозита
	depositEvent := account.Deposit(100)
	if depositEvent.Type != "deposit" {
		t.Errorf("Expected event type 'deposit', got %s", depositEvent.Type)
	}

	if depositEvent.Data != 100 {
		t.Errorf("Expected event data 100, got %v", depositEvent.Data)
	}

	// Создаем событие вывода
	withdrawEvent := account.Withdraw(50)
	if withdrawEvent.Type != "withdraw" {
		t.Errorf("Expected event type 'withdraw', got %s", withdrawEvent.Type)
	}

	if withdrawEvent.Data != 50 {
		t.Errorf("Expected event data 50, got %v", withdrawEvent.Data)
	}
}

func TestEventStoreConcurrent(t *testing.T) {
	store := NewEventStore()

	// Создаем события для разных агрегатов
	events1 := []Event{
		{ID: "1", Type: "deposit", Data: 100, Timestamp: time.Now().UnixMilli(), Version: 1},
	}

	events2 := []Event{
		{ID: "2", Type: "deposit", Data: 200, Timestamp: time.Now().UnixMilli(), Version: 1},
	}

	// Добавляем события для разных агрегатов
	err1 := store.AppendEvents("account1", events1)
	err2 := store.AppendEvents("account2", events2)

	if err1 != nil {
		t.Errorf("Expected no error for account1, got %v", err1)
	}

	if err2 != nil {
		t.Errorf("Expected no error for account2, got %v", err2)
	}

	// Проверяем версии агрегатов
	version1 := store.GetAggregateVersion("account1")
	version2 := store.GetAggregateVersion("account2")

	if version1 != 1 {
		t.Errorf("Expected version 1 for account1, got %d", version1)
	}

	if version2 != 1 {
		t.Errorf("Expected version 1 for account2, got %d", version2)
	}
}

