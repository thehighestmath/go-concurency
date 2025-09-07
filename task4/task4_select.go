package task4

import (
	"time"
)

// Task 4: Select Statement
// Реализуйте функцию, которая использует select для чтения из нескольких каналов.
// Функция должна читать числа из двух каналов и возвращать их сумму.

func sumCh(ch <-chan int, timeout int64) int {
	s := 0
	done := make(chan bool)
	res := make(chan int)
	go func() {
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		done <- true
	}()
	go func() {
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					res <- s
					return
				}
				s += v
			case <-done:
				res <- s
				return
			}
		}
	}()

	return <-res
}

// SumFromTwoChannels читает числа из двух каналов и возвращает их сумму
// Функция должна завершиться через timeout миллисекунд
func SumFromTwoChannels(ch1, ch2 <-chan int, timeout int64) int {
	s := 0
	ch := make(chan int)

	go func() {
		ch <- sumCh(ch1, timeout)
	}()
	go func() {
		ch <- sumCh(ch2, timeout)
	}()

	s += <-ch
	s += <-ch
	return s
}
