package task7

import (
	"context"
	"slices"
	"sync"
	"time"
)

// Task 7: Context Cancellation
// Реализуйте функцию, которая использует context для отмены долго выполняющихся операций.
// Функция должна запускать несколько горутин и отменять их через context.

// LongRunningTask представляет долго выполняющуюся задачу
type LongRunningTask struct {
	ID       int
	Duration int64 // Duration in milliseconds
}

func (t *LongRunningTask) Execute() string {
	time.Sleep(time.Duration(t.Duration) * time.Millisecond)
	return "completed"
}

// ProcessTasksWithContext обрабатывает задачи с возможностью отмены через context
// tasks - слайс задач для обработки
// timeout - максимальное время выполнения в миллисекундах
func ProcessTasksWithContext(tasks []LongRunningTask, timeout int64) []string {
	type result struct {
		ID     int
		Result string
	}
	out := make(chan result)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()
	go func() {
		var wg sync.WaitGroup
		for _, task := range tasks {
			wg.Add(1)
			go func(ctx context.Context, task LongRunningTask) {
				defer wg.Done()
				ch := make(chan string)
				go func() {
					ch <- task.Execute()
				}()
				select {
				case <-ctx.Done():
					out <- result{ID: task.ID, Result: "cancelled"}
				case x := <-ch:
					out <- result{ID: task.ID, Result: x}
				}
			}(ctx, task)
		}
		wg.Wait()
		close(out)
	}()
	arr := make([]result, 0)
	for x := range out {
		arr = append(arr, x)
	}
	slices.SortFunc(arr, func(l, r result) int { return r.ID - l.ID })
	out2 := make([]string, 0, len(arr))
	for _, x := range arr {
		out2 = append(out2, x.Result)
	}
	// TODO: Реализуйте эту функцию
	// 1. Создайте context с timeout
	// 2. Запустите горутину для каждой задачи
	// 3. Каждая горутина должна:
	//    - проверять context.Done() для отмены
	//    - симулировать работу в течение task.Duration
	//    - возвращать результат или ошибку отмены
	// 4. Соберите результаты от всех горутин
	// 5. Верните слайс результатов

	return out2
}
