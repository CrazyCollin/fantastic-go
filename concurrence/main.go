package main

import "fantastic-go/concurrence/concurrent"

func main() {
	//var c Counter
	//var wg sync.WaitGroup
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		defer wg.Done()
	//		for i := 0; i < 10000; i++ {
	//			c.Incr()
	//		}
	//	}()
	//}
	//wg.Wait()
	//fmt.Println(c.Count())

	//var c ConcurrenceCounter
	//var wg sync.WaitGroup
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//	go worker(&c, &wg)
	//}
	//wg.Wait()
	//println(c.Count())

	//condUse()

	concurrent.QueueUse()
}
