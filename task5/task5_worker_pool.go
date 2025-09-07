package task5

import (
	"sync"
)

// Task 5: Worker Pool Pattern
// Реализуйте паттерн Worker Pool. Создайте пул из N воркеров,
// которые обрабатывают задачи из очереди.

// "fmt"
// "sync"

// Task представляет задачу для обработки
type Task struct {
	ID   int
	Data string
}

// Result представляет результат обработки задачи
type Result struct {
	TaskID int
	Output string
	Error  error
}

// WorkerPool управляет пулом воркеров
type WorkerPool struct {
	resultChannel chan Result
	tasksChannel  chan Task
	workerCount   int
	wg            sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		resultChannel: make(chan Result),
		tasksChannel:  make(chan Task),
		workerCount:   numWorkers,
		wg:            sync.WaitGroup{},
	}
}

// Start запускает воркеров
func (wp *WorkerPool) Start() {
	wp.wg.Add(wp.workerCount)
	for i := 0; i < wp.workerCount; i++ {
		go func() {
			defer wp.wg.Done()
			for task := range wp.tasksChannel {
				wp.resultChannel <- Result{
					TaskID: task.ID,
					Output: task.Data,
				}
			}
		}()
	}
	// Запустите numWorkers горутин-воркеров
}

// SubmitTask отправляет задачу в пул
func (wp *WorkerPool) SubmitTask(task Task) {
	go func() {
		wp.tasksChannel <- task
	}()
}

// GetResult возвращает результат обработки
func (wp *WorkerPool) GetResult() <-chan Result {
	return wp.resultChannel
}

// Stop останавливает пул воркеров
func (wp *WorkerPool) Stop() {
	// Закройте канал задач и дождитесь завершения всех воркеров
	close(wp.tasksChannel)
	wp.wg.Wait()
}
