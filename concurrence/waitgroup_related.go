package main

import "sync"

type ConcurrenceCounter struct {
	mu    sync.Mutex
	count int64
}

func (c *ConcurrenceCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *ConcurrenceCounter) Count() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func worker(counter *ConcurrenceCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Inc()
}
