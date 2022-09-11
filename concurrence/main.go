package main

import (
	"fmt"
	"sync"
)

func main() {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				c.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.Count())
}
