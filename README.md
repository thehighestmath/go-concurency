# Go Concurrency Tasks

Набор задач по конкурентности в Go для изучения и практики работы с горутинами, каналами, мьютексами и другими примитивами синхронизации.

## Структура проекта

```
go-concurrency/
├── go.mod                           # Модуль Go
├── README.md                        # Этот файл
├── Makefile                         # Команды для запуска тестов
├── task1_goroutines.go             # Задача 1: Основы горутин
├── task1_goroutines_test.go        # Тесты для задачи 1
├── task2_channels.go               # Задача 2: Каналы
├── task2_channels_test.go          # Тесты для задачи 2
├── task3_mutex.go                  # Задача 3: Мьютексы
├── task3_mutex_test.go             # Тесты для задачи 3
├── task4_select.go                 # Задача 4: Select
├── task4_select_test.go            # Тесты для задачи 4
├── task5_worker_pool.go            # Задача 5: Worker Pool
├── task5_worker_pool_test.go       # Тесты для задачи 5
├── task6_fan_out_fan_in.go         # Задача 6: Fan-out Fan-in
├── task6_fan_out_fan_in_test.go    # Тесты для задачи 6
├── task7_context.go                # Задача 7: Context
├── task7_context_test.go           # Тесты для задачи 7
├── task8_pipeline.go               # Задача 8: Pipeline
├── task8_pipeline_test.go          # Тесты для задачи 8
├── task9_rate_limiter.go           # Задача 9: Rate Limiter
├── task9_rate_limiter_test.go      # Тесты для задачи 9
├── task10_barrier.go               # Задача 10: Barrier
├── task10_barrier_test.go          # Тесты для задачи 10
├── task11_semaphore.go             # Задача 11: Семафор
├── task11_semaphore_test.go        # Тесты для задачи 11
├── task12_race_condition.go        # Задача 12: Race Condition
├── task12_race_condition_test.go   # Тесты для задачи 12
├── task13_circuit_breaker.go       # Задача 13: Circuit Breaker
├── task13_circuit_breaker_test.go  # Тесты для задачи 13
├── task14_bulkhead.go              # Задача 14: Bulkhead Pattern
├── task14_bulkhead_test.go         # Тесты для задачи 14
├── task15_retry.go                 # Задача 15: Retry Pattern
├── task15_retry_test.go            # Тесты для задачи 15
├── task16_timeout.go               # Задача 16: Timeout Pattern
├── task16_timeout_test.go          # Тесты для задачи 16
├── task17_broadcast.go             # Задача 17: Broadcast Pattern
├── task17_broadcast_test.go        # Тесты для задачи 17
├── task18_priority_queue.go        # Задача 18: Priority Queue
├── task18_priority_queue_test.go   # Тесты для задачи 18
├── task19_cache.go                 # Задача 19: Concurrent Cache
├── task19_cache_test.go            # Тесты для задачи 19
├── task20_actor_model.go           # Задача 20: Actor Model
├── task20_actor_model_test.go      # Тесты для задачи 20
├── task21_event_sourcing.go        # Задача 21: Event Sourcing
├── task21_event_sourcing_test.go   # Тесты для задачи 21
├── task22_stream_processing.go     # Задача 22: Stream Processing
└── task22_stream_processing_test.go # Тесты для задачи 22
```

## Установка и запуск

1. Убедитесь, что у вас установлен Go версии 1.21 или выше
2. Клонируйте или скачайте проект
3. Установите зависимости:
   ```bash
   go mod tidy
   ```

## Запуск тестов

### Все тесты
```bash
make test
# или
go test ./...
```

### Конкретная задача
```bash
make test-task1
# или
go test -run TestRunNGoroutines
```

### С покрытием кода
```bash
make test-coverage
# или
go test -cover ./...
```

## Описание задач

### Задача 1: Основы горутин (task1_goroutines.go)
**Цель**: Изучить базовые принципы работы с горутинами и WaitGroup.

**Задание**: Реализуйте функцию `RunNGoroutines(n int)`, которая:
- Запускает N горутин
- Каждая горутина печатает свой номер и засыпает на случайное время (1-100ms)
- Функция дожидается завершения всех горутин

**Ключевые концепции**: `sync.WaitGroup`, `go` keyword, `time.Sleep()`

### Задача 2: Каналы (task2_channels.go)
**Цель**: Изучить основы работы с каналами для коммуникации между горутинами.

**Задание**: Реализуйте функцию `SumNumbersWithChannels(n int) int`, которая:
- Создает канал для передачи чисел
- Запускает горутину-производителя, отправляющую числа 1..n
- Запускает горутину-потребителя, суммирующую числа
- Возвращает сумму всех чисел

**Ключевые концепции**: `chan`, `<-`, `close()`, producer-consumer pattern

### Задача 3: Мьютексы (task3_mutex.go)
**Цель**: Изучить синхронизацию доступа к общим ресурсам с помощью мьютексов.

**Задание**: Реализуйте потокобезопасный счетчик `SafeCounter` с методами:
- `Increment()` - увеличивает значение на 1
- `Decrement()` - уменьшает значение на 1
- `GetValue()` - возвращает текущее значение

**Ключевые концепции**: `sync.Mutex`, `sync.RWMutex`, thread safety

### Задача 4: Select (task4_select.go)
**Цель**: Изучить использование `select` для работы с несколькими каналами.

**Задание**: Реализуйте функцию `SumFromTwoChannels()`, которая:
- Читает числа из двух каналов одновременно
- Суммирует все полученные числа
- Завершается по timeout
- Возвращает итоговую сумму

**Ключевые концепции**: `select`, `time.After()`, non-blocking operations

### Задача 5: Worker Pool (task5_worker_pool.go)
**Цель**: Изучить паттерн Worker Pool для обработки задач.

**Задание**: Реализуйте `WorkerPool` с методами:
- `Start()` - запускает воркеров
- `SubmitTask()` - отправляет задачу в пул
- `GetResult()` - возвращает канал с результатами
- `Stop()` - останавливает пул

**Ключевые концепции**: Worker pool pattern, task queue, result aggregation

### Задача 6: Fan-out Fan-in (task6_fan_out_fan_in.go)
**Цель**: Изучить паттерн Fan-out Fan-in для параллельной обработки данных.

**Задание**: Реализуйте функцию `FanOutFanIn()`, которая:
- Fan-out: распределяет числа между несколькими воркерами
- Каждый воркер применяет функцию обработки
- Fan-in: собирает результаты в один канал

**Ключевые концепции**: Fan-out Fan-in pattern, parallel processing

### Задача 7: Context (task7_context.go)
**Цель**: Изучить использование context для отмены операций.

**Задание**: Реализуйте функцию `ProcessTasksWithContext()`, которая:
- Обрабатывает список задач с возможностью отмены
- Использует context с timeout
- Возвращает результаты или информацию об отмене

**Ключевые концепции**: `context.Context`, `context.WithTimeout()`, cancellation

### Задача 8: Pipeline (task8_pipeline.go)
**Цель**: Изучить паттерн Pipeline для последовательной обработки данных.

**Задание**: Реализуйте pipeline из трех этапов:
1. `CreateNumberGenerator()` - генерирует числа 1..n
2. `FilterEvenNumbers()` - фильтрует четные числа
3. `DoubleNumbers()` - удваивает числа

**Ключевые концепции**: Pipeline pattern, data transformation

### Задача 9: Rate Limiter (task9_rate_limiter.go)
**Цель**: Изучить реализацию rate limiter с использованием token bucket.

**Задание**: Реализуйте `RateLimiter` с методами:
- `Allow()` - проверяет, можно ли выполнить запрос
- `Wait()` - ждет доступности токена
- `AcquireWithContext()` - получает токен с учетом контекста

**Ключевые концепции**: Rate limiting, token bucket algorithm, throttling

### Задача 10: Barrier (task10_barrier.go)
**Цель**: Изучить механизм синхронизации Barrier.

**Задание**: Реализуйте `Barrier`, который:
- Заставляет N горутин ждать друг друга
- Разрешает продолжение только когда все горутины достигли barrier
- Поддерживает повторное использование

**Ключевые концепции**: `sync.Cond`, barrier synchronization, rendezvous point

### Задача 11: Семафор (task11_semaphore.go)
**Цель**: Изучить реализацию семафора для ограничения ресурсов.

**Задание**: Реализуйте `Semaphore` с методами:
- `Acquire()` - получает токен (блокирующий)
- `TryAcquire()` - пытается получить токен (неблокирующий)
- `AcquireWithContext()` - получает токен с учетом контекста
- `Release()` - возвращает токен

**Ключевые концепции**: Semaphore, resource limiting, bounded concurrency

### Задача 12: Race Condition (task12_race_condition.go)
**Цель**: Найти и исправить race condition в коде.

**Задание**: В коде `BankAccount` есть race condition. Найдите и исправьте его, используя подходящие механизмы синхронизации. Также реализуйте безопасный перевод денег между счетами.

**Ключевые концепции**: Race condition detection, deadlock prevention, proper locking order

### Задача 13: Circuit Breaker (task13_circuit_breaker.go)
**Цель**: Изучить паттерн Circuit Breaker для защиты от каскадных сбоев.

**Задание**: Реализуйте `CircuitBreaker` с состояниями Closed, Open, Half-Open. Circuit breaker должен переключаться между состояниями на основе количества успешных/неуспешных вызовов.

**Ключевые концепции**: Circuit breaker pattern, fault tolerance, state machine

### Задача 14: Bulkhead Pattern (task14_bulkhead.go)
**Цель**: Изучить паттерн Bulkhead для изоляции ресурсов.

**Задание**: Реализуйте `BulkheadManager`, который создает отдельные пулы ресурсов для разных типов операций (чтение, запись, удаление). Каждый пул должен иметь свой лимит конкурентности.

**Ключевые концепции**: Resource isolation, fault isolation, bounded concurrency

### Задача 15: Retry Pattern (task15_retry.go)
**Цель**: Изучить паттерн Retry с экспоненциальной задержкой.

**Задание**: Реализуйте `RetryManager` с экспоненциальной задержкой, jitter и кастомной логикой retry. Поддержите настройки максимального количества попыток, базовой задержки и максимальной задержки.

**Ключевые концепции**: Retry pattern, exponential backoff, jitter, resilience

### Задача 16: Timeout Pattern (task16_timeout.go)
**Цель**: Изучить различные паттерны работы с таймаутами.

**Задание**: Реализуйте `TimeoutManager` для управления таймаутами и функции `ExecuteWithTimeout`/`ExecuteWithDeadline` для выполнения операций с ограничением времени.

**Ключевые концепции**: Timeout management, deadline handling, context cancellation

### Задача 17: Broadcast Pattern (task17_broadcast.go)
**Цель**: Изучить паттерн Broadcast для уведомления множества получателей.

**Задание**: Реализуйте `BroadcastChannel` и `BroadcastManager` для отправки сообщений множеству подписчиков. Поддержите подписку/отписку и изоляцию по темам.

**Ключевые концепции**: Publisher-subscriber pattern, message broadcasting, topic isolation

### Задача 18: Priority Queue (task18_priority_queue.go)
**Цель**: Изучить приоритетную очередь с поддержкой конкурентного доступа.

**Задание**: Реализуйте `PriorityQueue` и `ConcurrentPriorityQueue` с поддержкой блокирующих и неблокирующих операций. Обеспечьте правильный порядок элементов по приоритету.

**Ключевые концепции**: Priority queue, heap data structure, concurrent access

### Задача 19: Concurrent Cache (task19_cache.go)
**Цель**: Изучить реализацию потокобезопасного кэша с TTL.

**Задание**: Реализуйте `Cache` с TTL, автоматической очисткой и LRU политикой. Поддержите операции Set, Get, Delete с потокобезопасностью.

**Ключевые концепции**: Cache implementation, TTL, LRU eviction, thread safety

### Задача 20: Actor Model (task20_actor_model.go)
**Цель**: Изучить модель Actor для изоляции состояния.

**Задание**: Реализуйте `ActorSystem` и `Actor` для обмена сообщениями между изолированными акторами. Создайте `CounterActor` как пример использования.

**Ключевые концепции**: Actor model, message passing, state isolation, concurrency

### Задача 21: Event Sourcing (task21_event_sourcing.go)
**Цель**: Изучить паттерн Event Sourcing для хранения состояния.

**Задание**: Реализуйте `EventStore`, `EventHandler`, `EventBus` и `BankAccountAggregate` для демонстрации event sourcing. Поддержите версионирование и восстановление состояния.

**Ключевые концепции**: Event sourcing, event store, aggregate, domain events

### Задача 22: Stream Processing (task22_stream_processing.go)
**Цель**: Изучить систему обработки потоков данных.

**Задание**: Реализуйте `StreamProcessor`, `StreamAggregator`, `StreamJoiner`, `StreamFilter` и `StreamPipeline` для обработки потоков данных с окнами и агрегацией.

**Ключевые концепции**: Stream processing, windowing, aggregation, data pipelines

## Советы по решению

1. **Начните с простых задач** (1-4) для понимания основ
2. **Изучайте тесты** - они показывают ожидаемое поведение
3. **Используйте `go test -race`** для обнаружения race conditions
4. **Читайте документацию** Go по конкурентности
5. **Экспериментируйте** с разными подходами

## Полезные ресурсы

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Go by Example - Select](https://gobyexample.com/select)

## Проверка решений

После реализации каждой задачи запустите тесты:

```bash
# Проверка конкретной задачи
go test -v -run TestRunNGoroutines

# Проверка всех тестов
go test -v ./...

# Проверка на race conditions
go test -race ./...
```

Удачи в изучении конкурентности в Go! 🚀

```
# Установить зависимости
make install

# Запустить все тесты
make test

# Запустить тесты конкретной задачи
make test-task1

# Запустить с проверкой race conditions
make test-race

# Посмотреть справку
make help
```