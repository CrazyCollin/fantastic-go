package main

import (
	"sync"
)

type Counter struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count int64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Count() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
