// package main

// import (
// 	"sync/atomic"
// )

// // SafeCounter - потокобезопасный счетчик
// type SafeCounter struct {
// 	v atomic.Int32
// }

// func NewSafeCounter() *SafeCounter {
// 	return &SafeCounter{}
// }

// func (c *SafeCounter) Increment() {
// 	c.v.Add(1)
// }

// func (c *SafeCounter) Decrement() {
// 	c.v.Add(-1)
// }

// func (c *SafeCounter) GetValue() int {
// 	return int(c.v.Load())
// }

package task3

import (
	"sync"
)

// SafeCounter - потокобезопасный счетчик
type SafeCounter struct {
	m sync.Mutex
	v int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (c *SafeCounter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.v++
}

func (c *SafeCounter) Decrement() {
	c.m.Lock()
	defer c.m.Unlock()
	c.v--
}

func (c *SafeCounter) GetValue() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.v
}
