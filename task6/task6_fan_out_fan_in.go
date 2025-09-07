package task6

import (
	"sync"
)

func merge(ins []chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, ch := range ins {
		go func(ch <-chan int) {
			defer wg.Done()
			for x := range ch {
				out <- x
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Task 6: Fan-out Fan-in Pattern
// Реализуйте паттерн Fan-out Fan-in. Создайте функцию, которая:
// 1. Fan-out: распределяет числа из одного канала между несколькими воркерами
// 2. Fan-in: собирает результаты от всех воркеров в один канал

// FanOutFanIn распределяет числа между воркерами и собирает результаты
// numbers - канал с входными числами
// numWorkers - количество воркеров
// processFunc - функция обработки числа
func FanOutFanIn(numbers <-chan int, numWorkers int, processFunc func(int) int) <-chan int {
	chs := make([]chan int, 0)
	for i := 0; i < numWorkers; i++ {
		chs = append(chs, make(chan int))
		go func(i int) {
			for {
				x, ok := <-numbers
				if !ok {
					close(chs[i])
					return
				}
				chs[i] <- processFunc(x)
			}
		}(i)
	}

	out := make(chan chan int)
	go func() {
		x := merge(chs)
		out <- x
	}()

	return <-out

	// TODO: Реализуйте эту функцию
	// 1. Fan-out: создайте numWorkers воркеров, каждый читает из numbers
	// 2. Каждый воркер применяет processFunc к числам
	// 3. Fan-in: соберите результаты от всех воркеров в один канал
}
