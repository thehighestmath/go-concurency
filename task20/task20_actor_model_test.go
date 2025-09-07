package task20

import (
	"sync"
	"testing"
	"time"
)

func TestActorSystemBasic(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора
	actor := system.SpawnActor("counter", 0, func(state interface{}, msg Message) interface{} {
		count := state.(int)
		switch msg.Type {
		case "increment":
			return count + 1
		case "decrement":
			return count - 1
		default:
			return count
		}
	})

	if actor == nil {
		t.Error("Actor should be created")
	}

	// Проверяем количество акторов
	if system.GetActorCount() != 1 {
		t.Errorf("Expected 1 actor, got %d", system.GetActorCount())
	}
}

func TestActorSystemSendMessage(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора
	system.SpawnActor("counter", 0, func(state interface{}, msg Message) interface{} {
		count := state.(int)
		switch msg.Type {
		case "increment":
			return count + 1
		case "get":
			if msg.ReplyTo != nil {
				msg.ReplyTo <- count
			}
			return count
		default:
			return count
		}
	})

	// Отправляем сообщение
	replyChan := make(chan interface{}, 1)
	msg := Message{
		Type:    "increment",
		Data:    nil,
		ReplyTo: nil,
	}
	system.SendMessage("counter", msg)

	// Отправляем сообщение с ответом
	msg = Message{
		Type:    "get",
		Data:    nil,
		ReplyTo: replyChan,
	}
	system.SendMessage("counter", msg)

	// Проверяем ответ
	select {
	case count := <-replyChan:
		if count != 1 {
			t.Errorf("Expected count 1, got %v", count)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive reply")
	}
}

func TestActorSystemSendMessageAsync(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора
	system.SpawnActor("counter", 0, func(state interface{}, msg Message) interface{} {
		count := state.(int)
		switch msg.Type {
		case "increment":
			return count + 1
		default:
			return count
		}
	})

	// Отправляем асинхронное сообщение
	system.SendMessageAsync("counter", "increment", nil)

	// Ждем обработки
	time.Sleep(10 * time.Millisecond)

	// Проверяем, что сообщение обработано (нет ошибок)
	// В реальной реализации здесь можно было бы проверить состояние
}

func TestActorSystemSendMessageSync(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора
	system.SpawnActor("counter", 0, func(state interface{}, msg Message) interface{} {
		count := state.(int)
		switch msg.Type {
		case "increment":
			return count + 1
		case "get":
			return count
		default:
			return count
		}
	})

	// Отправляем синхронное сообщение
	result := system.SendMessageSync("counter", "increment", nil)
	if result != 1 {
		t.Errorf("Expected result 1, got %v", result)
	}

	// Отправляем еще одно сообщение
	result = system.SendMessageSync("counter", "get", nil)
	if result != 1 {
		t.Errorf("Expected result 1, got %v", result)
	}
}

func TestActorSystemStopActor(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора
	system.SpawnActor("counter", 0, func(state interface{}, msg Message) interface{} {
		return state
	})

	// Проверяем, что актор создан
	if system.GetActorCount() != 1 {
		t.Errorf("Expected 1 actor, got %d", system.GetActorCount())
	}

	// Останавливаем актора
	system.StopActor("counter")

	// Проверяем, что актор остановлен
	if system.GetActorCount() != 0 {
		t.Errorf("Expected 0 actors, got %d", system.GetActorCount())
	}
}

func TestActorSystemStopAllActors(t *testing.T) {
	system := NewActorSystem()

	// Создаем несколько акторов
	system.SpawnActor("actor1", 0, func(state interface{}, msg Message) interface{} {
		return state
	})
	system.SpawnActor("actor2", 0, func(state interface{}, msg Message) interface{} {
		return state
	})

	// Проверяем, что акторы созданы
	if system.GetActorCount() != 2 {
		t.Errorf("Expected 2 actors, got %d", system.GetActorCount())
	}

	// Останавливаем всех акторов
	system.StopAllActors()

	// Проверяем, что все акторы остановлены
	if system.GetActorCount() != 0 {
		t.Errorf("Expected 0 actors, got %d", system.GetActorCount())
	}
}

func TestCounterActorBasic(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора-счетчик
	counter := NewCounterActor(system, "counter")

	if counter == nil {
		t.Error("Counter actor should be created")
	}

	// Проверяем начальное значение
	value := counter.GetValue()
	if value != 0 {
		t.Errorf("Expected initial value 0, got %d", value)
	}
}

func TestCounterActorIncrement(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора-счетчик
	counter := NewCounterActor(system, "counter")

	// Увеличиваем счетчик
	counter.Increment()

	// Проверяем значение
	value := counter.GetValue()
	if value != 1 {
		t.Errorf("Expected value 1, got %d", value)
	}

	// Увеличиваем еще раз
	counter.Increment()

	// Проверяем значение
	value = counter.GetValue()
	if value != 2 {
		t.Errorf("Expected value 2, got %d", value)
	}
}

func TestCounterActorDecrement(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора-счетчик
	counter := NewCounterActor(system, "counter")

	// Увеличиваем счетчик
	counter.Increment()
	counter.Increment()

	// Уменьшаем счетчик
	counter.Decrement()

	// Проверяем значение
	value := counter.GetValue()
	if value != 1 {
		t.Errorf("Expected value 1, got %d", value)
	}
}

func TestCounterActorReset(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора-счетчик
	counter := NewCounterActor(system, "counter")

	// Увеличиваем счетчик
	counter.Increment()
	counter.Increment()

	// Проверяем значение
	value := counter.GetValue()
	if value != 2 {
		t.Errorf("Expected value 2, got %d", value)
	}

	// Сбрасываем счетчик
	counter.Reset()

	// Проверяем значение
	value = counter.GetValue()
	if value != 0 {
		t.Errorf("Expected value 0 after reset, got %d", value)
	}
}

func TestActorSystemConcurrent(t *testing.T) {
	system := NewActorSystem()

	// Создаем актора-счетчик
	counter := NewCounterActor(system, "counter")

	var wg sync.WaitGroup

	// Запускаем горутины, которые увеличивают счетчик
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	// Проверяем финальное значение
	value := counter.GetValue()
	if value != 10 {
		t.Errorf("Expected value 10, got %d", value)
	}
}

func TestActorSystemMultipleActors(t *testing.T) {
	system := NewActorSystem()

	// Создаем несколько акторов-счетчиков
	counter1 := NewCounterActor(system, "counter1")
	counter2 := NewCounterActor(system, "counter2")

	// Проверяем, что оба актора созданы
	if system.GetActorCount() != 2 {
		t.Errorf("Expected 2 actors, got %d", system.GetActorCount())
	}

	// Увеличиваем счетчики
	counter1.Increment()
	counter2.Increment()
	counter2.Increment()

	// Проверяем значения
	if counter1.GetValue() != 1 {
		t.Errorf("Counter1 expected 1, got %d", counter1.GetValue())
	}
	if counter2.GetValue() != 2 {
		t.Errorf("Counter2 expected 2, got %d", counter2.GetValue())
	}
}

