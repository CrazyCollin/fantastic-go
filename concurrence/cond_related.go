package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type FIFO struct {
	mu    sync.Mutex
	cond  *sync.Cond
	queue []interface{}
}

func (f *FIFO) Offer(item interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.queue = append(f.queue, item)
	f.cond.Broadcast()
	return nil
}

func (f *FIFO) Pop() interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()
	for {
		for len(f.queue) == 0 {
			f.cond.Wait()
		}
		item := f.queue[0]
		f.queue = f.queue[1:]
		return item
	}
}

func queueUse() {
	l := sync.Mutex{}
	fifo := &FIFO{
		mu:    l,
		cond:  sync.NewCond(&l),
		queue: []interface{}{},
	}
	go func() {
		for {
			_ = fifo.Offer(rand.Int())
		}
	}()
	time.Sleep(time.Second)
	go func() {
		for {
			fmt.Printf("goroutine 1 pop %v \n", fifo.Pop())
		}
	}()
	go func() {
		for {
			fmt.Printf("goroutine 2 pop %v\n", fifo.Pop())
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func condUse() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("goroutine %d is ready", i)
			c.Broadcast()
		}(i)
	}
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("goroutine weak up")
	}
	c.L.Unlock()

	log.Println("all goroutine is ready")
}
