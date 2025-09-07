package task16

import (
	"testing"
	"time"
)

func TestTimeoutManagerBasic(t *testing.T) {
	tm := NewTimeoutManager()

	callbackCalled := false
	config := TimeoutConfig{
		Duration: 50,
		Callback: func() {
			callbackCalled = true
		},
		ID: "test1",
	}

	tm.SetTimeout(config)

	// Ждем выполнения callback
	time.Sleep(60 * time.Millisecond)

	if !callbackCalled {
		t.Error("Callback should be called after timeout")
	}
}

func TestTimeoutManagerClearTimeout(t *testing.T) {
	tm := NewTimeoutManager()

	callbackCalled := false
	config := TimeoutConfig{
		Duration: 100,
		Callback: func() {
			callbackCalled = true
		},
		ID: "test1",
	}

	tm.SetTimeout(config)

	// Отменяем таймаут
	tm.ClearTimeout("test1")

	// Ждем дольше, чем таймаут
	time.Sleep(120 * time.Millisecond)

	if callbackCalled {
		t.Error("Callback should not be called after clearing timeout")
	}
}

func TestTimeoutManagerMultipleTimeouts(t *testing.T) {
	tm := NewTimeoutManager()

	callbacks := make(map[string]bool)

	// Устанавливаем несколько таймаутов
	for i := 0; i < 3; i++ {
		id := "test" + string(rune(i))
		config := TimeoutConfig{
			Duration: 50,
			Callback: func() {
				callbacks[id] = true
			},
			ID: id,
		}
		tm.SetTimeout(config)
	}

	// Ждем выполнения всех callback'ов
	time.Sleep(60 * time.Millisecond)

	if len(callbacks) != 3 {
		t.Errorf("Expected 3 callbacks, got %d", len(callbacks))
	}
}

func TestTimeoutManagerClearAllTimeouts(t *testing.T) {
	tm := NewTimeoutManager()

	callbackCalled := false
	config := TimeoutConfig{
		Duration: 100,
		Callback: func() {
			callbackCalled = true
		},
		ID: "test1",
	}

	tm.SetTimeout(config)

	// Отменяем все таймауты
	tm.ClearAllTimeouts()

	// Ждем дольше, чем таймаут
	time.Sleep(120 * time.Millisecond)

	if callbackCalled {
		t.Error("Callback should not be called after clearing all timeouts")
	}
}

func TestTimeoutManagerActiveTimeouts(t *testing.T) {
	tm := NewTimeoutManager()

	// Проверяем начальное состояние
	if tm.GetActiveTimeouts() != 0 {
		t.Errorf("Expected 0 active timeouts, got %d", tm.GetActiveTimeouts())
	}

	// Устанавливаем таймаут
	config := TimeoutConfig{
		Duration: 100,
		Callback: func() {},
		ID:       "test1",
	}
	tm.SetTimeout(config)

	// Проверяем, что таймаут активен
	if tm.GetActiveTimeouts() != 1 {
		t.Errorf("Expected 1 active timeout, got %d", tm.GetActiveTimeouts())
	}

	// Отменяем таймаут
	tm.ClearTimeout("test1")

	// Проверяем, что таймаут неактивен
	if tm.GetActiveTimeouts() != 0 {
		t.Errorf("Expected 0 active timeouts after clear, got %d", tm.GetActiveTimeouts())
	}
}

func TestExecuteWithTimeout(t *testing.T) {
	// Тестируем успешное выполнение
	err := ExecuteWithTimeout(func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}, 50)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
}

func TestExecuteWithTimeoutExpired(t *testing.T) {
	// Тестируем таймаут
	err := ExecuteWithTimeout(func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}, 50)

	if err == nil {
		t.Error("Expected timeout error, got success")
	}
}

func TestExecuteWithTimeoutError(t *testing.T) {
	// Тестируем ошибку в функции
	err := ExecuteWithTimeout(func() error {
		return &TestError{message: "test error"}
	}, 50)

	if err == nil {
		t.Error("Expected error, got success")
	}
}

func TestExecuteWithDeadline(t *testing.T) {
	// Тестируем успешное выполнение с дедлайном
	deadline := time.Now().Add(50 * time.Millisecond).UnixMilli()

	err := ExecuteWithDeadline(func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}, deadline)

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
}

func TestExecuteWithDeadlineExpired(t *testing.T) {
	// Тестируем истечение дедлайна
	deadline := time.Now().Add(10 * time.Millisecond).UnixMilli()

	err := ExecuteWithDeadline(func() error {
		time.Sleep(50 * time.Millisecond)
		return nil
	}, deadline)

	if err == nil {
		t.Error("Expected deadline error, got success")
	}
}

func TestTimeoutManagerConcurrent(t *testing.T) {
	tm := NewTimeoutManager()

	// Устанавливаем много таймаутов одновременно
	for i := 0; i < 10; i++ {
		go func(id int) {
			config := TimeoutConfig{
				Duration: 50,
				Callback: func() {},
				ID:       "test" + string(rune(id)),
			}
			tm.SetTimeout(config)
		}(i)
	}

	// Ждем установки всех таймаутов
	time.Sleep(10 * time.Millisecond)

	// Проверяем количество активных таймаутов
	active := tm.GetActiveTimeouts()
	if active != 10 {
		t.Errorf("Expected 10 active timeouts, got %d", active)
	}

	// Ждем истечения всех таймаутов
	time.Sleep(60 * time.Millisecond)

	// Проверяем, что все таймауты истекли
	active = tm.GetActiveTimeouts()
	if active != 0 {
		t.Errorf("Expected 0 active timeouts after expiration, got %d", active)
	}
}

