package task17

import (
	"sync"
	"testing"
	"time"
)

func TestBroadcastChannelBasic(t *testing.T) {
	bc := NewBroadcastChannel()

	// Подписываемся на broadcast
	sub1 := bc.Subscribe("sub1")
	sub2 := bc.Subscribe("sub2")

	// Проверяем количество подписчиков
	if bc.GetSubscriberCount() != 2 {
		t.Errorf("Expected 2 subscribers, got %d", bc.GetSubscriberCount())
	}

	// Отправляем сообщение
	bc.Broadcast("test message")

	// Проверяем получение сообщений
	select {
	case msg := <-sub1.Ch:
		if msg != "test message" {
			t.Errorf("Expected 'test message', got %v", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Subscriber 1 should receive message")
	}

	select {
	case msg := <-sub2.Ch:
		if msg != "test message" {
			t.Errorf("Expected 'test message', got %v", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Subscriber 2 should receive message")
	}
}

func TestBroadcastChannelUnsubscribe(t *testing.T) {
	bc := NewBroadcastChannel()

	// Подписываемся
	sub1 := bc.Subscribe("sub1")
	sub2 := bc.Subscribe("sub2")

	// Отписываемся
	bc.Unsubscribe("sub1")

	// Проверяем количество подписчиков
	if bc.GetSubscriberCount() != 1 {
		t.Errorf("Expected 1 subscriber, got %d", bc.GetSubscriberCount())
	}

	// Отправляем сообщение
	bc.Broadcast("test message")

	// Проверяем, что sub1 не получил сообщение
	select {
	case <-sub1.Ch:
		t.Error("Unsubscribed subscriber should not receive message")
	case <-time.After(50 * time.Millisecond):
		// Ожидаемое поведение
	}

	// Проверяем, что sub2 получил сообщение
	select {
	case msg := <-sub2.Ch:
		if msg != "test message" {
			t.Errorf("Expected 'test message', got %v", msg)
		}
	case <-time.After(50 * time.Millisecond):
		t.Error("Subscriber 2 should receive message")
	}
}

func TestBroadcastChannelMultipleMessages(t *testing.T) {
	bc := NewBroadcastChannel()

	sub := bc.Subscribe("sub1")

	// Отправляем несколько сообщений
	messages := []string{"msg1", "msg2", "msg3"}
	for _, msg := range messages {
		bc.Broadcast(msg)
	}

	// Проверяем получение всех сообщений
	for i, expectedMsg := range messages {
		select {
		case msg := <-sub.Ch:
			if msg != expectedMsg {
				t.Errorf("Message %d: expected '%s', got %v", i, expectedMsg, msg)
			}
		case <-time.After(100 * time.Millisecond):
			t.Errorf("Should receive message %d", i)
		}
	}
}

func TestBroadcastChannelClose(t *testing.T) {
	bc := NewBroadcastChannel()

	sub1 := bc.Subscribe("sub1")
	sub2 := bc.Subscribe("sub2")

	// Закрываем broadcast канал
	bc.Close()

	// Проверяем, что каналы подписчиков закрыты
	select {
	case _, ok := <-sub1.Ch:
		if ok {
			t.Error("Subscriber 1 channel should be closed")
		}
	case <-time.After(50 * time.Millisecond):
		t.Error("Subscriber 1 channel should be closed")
	}

	select {
	case _, ok := <-sub2.Ch:
		if ok {
			t.Error("Subscriber 2 channel should be closed")
		}
	case <-time.After(50 * time.Millisecond):
		t.Error("Subscriber 2 channel should be closed")
	}
}

func TestBroadcastManagerBasic(t *testing.T) {
	bm := NewBroadcastManager()

	// Получаем канал по теме
	topic1 := bm.GetChannel("topic1")
	topic2 := bm.GetChannel("topic2")

	// Подписываемся на разные темы
	sub1 := topic1.Subscribe("sub1")
	sub2 := topic2.Subscribe("sub2")

	// Отправляем сообщения в разные темы
	bm.BroadcastToTopic("topic1", "message1")
	bm.BroadcastToTopic("topic2", "message2")

	// Проверяем получение сообщений
	select {
	case msg := <-sub1.Ch:
		if msg != "message1" {
			t.Errorf("Expected 'message1', got %v", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Subscriber 1 should receive message")
	}

	select {
	case msg := <-sub2.Ch:
		if msg != "message2" {
			t.Errorf("Expected 'message2', got %v", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Subscriber 2 should receive message")
	}
}

func TestBroadcastManagerTopicIsolation(t *testing.T) {
	bm := NewBroadcastManager()

	// Получаем канал по теме
	topic1 := bm.GetChannel("topic1")

	// Подписываемся на тему
	sub := topic1.Subscribe("sub1")

	// Отправляем сообщение в другую тему
	bm.BroadcastToTopic("topic2", "message")

	// Проверяем, что сообщение не получено
	select {
	case <-sub.Ch:
		t.Error("Subscriber should not receive message from different topic")
	case <-time.After(50 * time.Millisecond):
		// Ожидаемое поведение
	}
}

func TestBroadcastChannelConcurrent(t *testing.T) {
	bc := NewBroadcastChannel()

	// Создаем много подписчиков
	var wg sync.WaitGroup
	subscribers := make([]*Subscriber, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			subscribers[index] = bc.Subscribe("sub" + string(rune(index)))
		}(i)
	}

	wg.Wait()

	// Проверяем количество подписчиков
	if bc.GetSubscriberCount() != 10 {
		t.Errorf("Expected 10 subscribers, got %d", bc.GetSubscriberCount())
	}

	// Отправляем сообщение
	bc.Broadcast("concurrent message")

	// Проверяем, что все подписчики получили сообщение
	for i, sub := range subscribers {
		select {
		case msg := <-sub.Ch:
			if msg != "concurrent message" {
				t.Errorf("Subscriber %d: expected 'concurrent message', got %v", i, msg)
			}
		case <-time.After(100 * time.Millisecond):
			t.Errorf("Subscriber %d should receive message", i)
		}
	}
}

func TestBroadcastChannelBlockedSubscriber(t *testing.T) {
	bc := NewBroadcastChannel()

	// Создаем подписчика с буферизованным каналом
	sub := bc.Subscribe("sub1")

	// Отправляем много сообщений быстро
	for i := 0; i < 5; i++ {
		bc.Broadcast("message" + string(rune(i)))
	}

	// Проверяем, что подписчик получает сообщения
	receivedCount := 0
	timeout := time.After(200 * time.Millisecond)

	for {
		select {
		case msg := <-sub.Ch:
			receivedCount++
			if msg == nil {
				t.Error("Received nil message")
			}
		case <-timeout:
			break
		}
	}

	// Должны получить хотя бы одно сообщение
	if receivedCount == 0 {
		t.Error("Should receive at least one message")
	}
}

