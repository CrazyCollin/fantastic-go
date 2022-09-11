package main

import "sync"

type SliceQueue struct {
	mu   sync.Mutex
	data []interface{}
}

func (q *SliceQueue) EnQueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, item)
}

func (q *SliceQueue) DeQueue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	res := q.data[0]
	q.data = q.data[:len(q.data)-1]
	q.mu.Unlock()
	return res
}
