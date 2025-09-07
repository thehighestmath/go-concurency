package task1

import (
	"math/rand"
	"sync"
	"time"
)

// RunNGoroutines запускает N горутин и ждет их завершения
func RunNGoroutines(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d := time.Duration(100*rand.Float64()) * time.Millisecond
			time.Sleep(d)
		}()
	}
	wg.Wait()
}
