package task22

import (
	"sync"
	"testing"
	"time"
)

func TestStreamProcessorBasic(t *testing.T) {
	config := WindowConfig{
		Size:     100,
		Slide:    50,
		Function: "sum",
	}

	processor := NewStreamProcessor(config)

	// Создаем каналы
	input := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем обработку
	go processor.ProcessData(input, output)

	// Отправляем данные
	input <- DataPoint{Timestamp: 1000, Value: 10, ID: "1"}
	input <- DataPoint{Timestamp: 1100, Value: 20, ID: "2"}
	input <- DataPoint{Timestamp: 1200, Value: 30, ID: "3"}

	// Закрываем входной канал
	close(input)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 60 { // 10 + 20 + 30
			t.Errorf("Expected sum 60, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive processed data")
	}
}

func TestStreamProcessorAverage(t *testing.T) {
	config := WindowConfig{
		Size:     100,
		Slide:    50,
		Function: "avg",
	}

	processor := NewStreamProcessor(config)

	// Создаем каналы
	input := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем обработку
	go processor.ProcessData(input, output)

	// Отправляем данные
	input <- DataPoint{Timestamp: 1000, Value: 10, ID: "1"}
	input <- DataPoint{Timestamp: 1100, Value: 20, ID: "2"}
	input <- DataPoint{Timestamp: 1200, Value: 30, ID: "3"}

	// Закрываем входной канал
	close(input)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 20 { // (10 + 20 + 30) / 3
			t.Errorf("Expected average 20, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive processed data")
	}
}

func TestStreamProcessorMax(t *testing.T) {
	config := WindowConfig{
		Size:     100,
		Slide:    50,
		Function: "max",
	}

	processor := NewStreamProcessor(config)

	// Создаем каналы
	input := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем обработку
	go processor.ProcessData(input, output)

	// Отправляем данные
	input <- DataPoint{Timestamp: 1000, Value: 10, ID: "1"}
	input <- DataPoint{Timestamp: 1100, Value: 30, ID: "2"}
	input <- DataPoint{Timestamp: 1200, Value: 20, ID: "3"}

	// Закрываем входной канал
	close(input)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 30 {
			t.Errorf("Expected max 30, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive processed data")
	}
}

func TestStreamProcessorCount(t *testing.T) {
	config := WindowConfig{
		Size:     100,
		Slide:    50,
		Function: "count",
	}

	processor := NewStreamProcessor(config)

	// Создаем каналы
	input := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем обработку
	go processor.ProcessData(input, output)

	// Отправляем данные
	input <- DataPoint{Timestamp: 1000, Value: 10, ID: "1"}
	input <- DataPoint{Timestamp: 1100, Value: 20, ID: "2"}
	input <- DataPoint{Timestamp: 1200, Value: 30, ID: "3"}

	// Закрываем входной канал
	close(input)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 3 {
			t.Errorf("Expected count 3, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive processed data")
	}
}

func TestStreamAggregatorBasic(t *testing.T) {
	aggregator := NewStreamAggregator()

	// Агрегируем данные
	aggregator.AggregateData("key1", 10)
	aggregator.AggregateData("key1", 20)
	aggregator.AggregateData("key2", 30)

	// Проверяем агрегат для key1
	stats1 := aggregator.GetAggregate("key1")
	if stats1["sum"] != 30 {
		t.Errorf("Expected sum 30 for key1, got %f", stats1["sum"])
	}

	if stats1["count"] != 2 {
		t.Errorf("Expected count 2 for key1, got %f", stats1["count"])
	}

	// Проверяем агрегат для key2
	stats2 := aggregator.GetAggregate("key2")
	if stats2["sum"] != 30 {
		t.Errorf("Expected sum 30 for key2, got %f", stats2["sum"])
	}

	if stats2["count"] != 1 {
		t.Errorf("Expected count 1 for key2, got %f", stats2["count"])
	}
}

func TestStreamAggregatorAllAggregates(t *testing.T) {
	aggregator := NewStreamAggregator()

	// Агрегируем данные
	aggregator.AggregateData("key1", 10)
	aggregator.AggregateData("key2", 20)

	// Получаем все агрегаты
	allStats := aggregator.GetAllAggregates()

	if len(allStats) != 2 {
		t.Errorf("Expected 2 aggregates, got %d", len(allStats))
	}

	if allStats["key1"]["sum"] != 10 {
		t.Errorf("Expected sum 10 for key1, got %f", allStats["key1"]["sum"])
	}

	if allStats["key2"]["sum"] != 20 {
		t.Errorf("Expected sum 20 for key2, got %f", allStats["key2"]["sum"])
	}
}

func TestStreamJoinerBasic(t *testing.T) {
	config := JoinConfig{
		WindowSize: 100,
		JoinType:   "inner",
	}

	joiner := NewStreamJoiner(config)

	// Создаем каналы
	stream1 := make(chan DataPoint, 10)
	stream2 := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем соединение
	go joiner.JoinStreams(stream1, stream2, output)

	// Отправляем данные в первый поток
	stream1 <- DataPoint{Timestamp: 1000, Value: 10, ID: "1"}

	// Отправляем данные во второй поток
	stream2 <- DataPoint{Timestamp: 1000, Value: 20, ID: "2"}

	// Закрываем каналы
	close(stream1)
	close(stream2)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 30 { // 10 + 20
			t.Errorf("Expected joined value 30, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive joined data")
	}
}

func TestStreamFilterBasic(t *testing.T) {
	config := FilterConfig{
		MinValue: 10,
		MaxValue: 50,
		Function: nil,
	}

	filter := NewStreamFilter(config)

	// Создаем каналы
	input := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем фильтрацию
	go filter.FilterStream(input, output)

	// Отправляем данные
	input <- DataPoint{Timestamp: 1000, Value: 5, ID: "1"}  // Должно быть отфильтровано
	input <- DataPoint{Timestamp: 1100, Value: 25, ID: "2"} // Должно пройти
	input <- DataPoint{Timestamp: 1200, Value: 60, ID: "3"} // Должно быть отфильтровано

	// Закрываем входной канал
	close(input)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 25 {
			t.Errorf("Expected filtered value 25, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive filtered data")
	}
}

func TestStreamFilterCustomFunction(t *testing.T) {
	config := FilterConfig{
		MinValue: 0,
		MaxValue: 0,
		Function: func(point DataPoint) bool {
			return point.Value > 20
		},
	}

	filter := NewStreamFilter(config)

	// Создаем каналы
	input := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем фильтрацию
	go filter.FilterStream(input, output)

	// Отправляем данные
	input <- DataPoint{Timestamp: 1000, Value: 10, ID: "1"} // Должно быть отфильтровано
	input <- DataPoint{Timestamp: 1100, Value: 30, ID: "2"} // Должно пройти

	// Закрываем входной канал
	close(input)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 30 {
			t.Errorf("Expected filtered value 30, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive filtered data")
	}
}

func TestStreamPipelineBasic(t *testing.T) {
	pipeline := NewStreamPipeline()

	// Создаем каналы
	input := make(chan DataPoint, 10)
	output := make(chan DataPoint, 10)

	// Запускаем пайплайн
	pipeline.Start(input, output)
	defer pipeline.Stop()

	// Отправляем данные
	input <- DataPoint{Timestamp: 1000, Value: 10, ID: "1"}
	input <- DataPoint{Timestamp: 1100, Value: 20, ID: "2"}

	// Закрываем входной канал
	close(input)

	// Ждем обработки
	time.Sleep(50 * time.Millisecond)

	// Проверяем результат
	select {
	case result := <-output:
		if result.Value != 30 { // 10 + 20
			t.Errorf("Expected processed value 30, got %f", result.Value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Should receive processed data")
	}
}

func TestStreamProcessorConcurrent(t *testing.T) {
	config := WindowConfig{
		Size:     100,
		Slide:    50,
		Function: "sum",
	}

	processor := NewStreamProcessor(config)

	// Создаем каналы
	input := make(chan DataPoint, 100)
	output := make(chan DataPoint, 100)

	// Запускаем обработку
	go processor.ProcessData(input, output)

	// Отправляем много данных
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			input <- DataPoint{Timestamp: int64(1000 + index), Value: float64(index), ID: string(rune(index))}
		}(i)
	}

	wg.Wait()
	close(input)

	// Ждем обработки
	time.Sleep(100 * time.Millisecond)

	// Проверяем, что получили результат
	select {
	case result := <-output:
		if result.Value != 45 { // 0 + 1 + 2 + ... + 9
			t.Errorf("Expected sum 45, got %f", result.Value)
		}
	case <-time.After(200 * time.Millisecond):
		t.Error("Should receive processed data")
	}
}

