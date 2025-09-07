package task8

// Task 8: Pipeline Pattern
// Реализуйте pipeline из трех этапов:
// 1. Генератор чисел
// 2. Фильтр четных чисел
// 3. Удвоитель чисел

// "fmt"

// PipelineStage представляет этап pipeline
type PipelineStage func(<-chan int) <-chan int

func CreateSliceFromChannel[T any](ch <-chan T) []T {
	out := make([]T, 0)
	for x := range ch {
		out = append(out, x)
	}
	return out
}

// CreateNumberGenerator создает генератор чисел от 1 до n
func CreateNumberGenerator(n int) <-chan int {
	if n < 0 {
		ch := make(chan int, 0)
		close(ch)
		return ch
	}
	ch := make(chan int, n)

	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// FilterEvenNumbers фильтрует четные числа
func FilterEvenNumbers(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range input {
			if x%2 == 0 {
				out <- x
			}
		}
		close(out)
	}()
	return out
}

// DoubleNumbers удваивает числа
func DoubleNumbers(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range input {
			out <- x * 2
		}
		close(out)
	}()
	return out
}

// RunPipeline запускает pipeline из трех этапов
func RunPipeline(n int) []int {
	gen := CreateNumberGenerator(n)
	even := FilterEvenNumbers(gen)
	double := DoubleNumbers(even)
	sl := CreateSliceFromChannel(double)
	return sl
}
