package concurrence

import (
	"sync"
)

var Wg = &sync.WaitGroup{}

type GoroutineLimit struct {
	num int
	ch  chan struct{}
}

func NewGoroutineLimit(num int) *GoroutineLimit {
	return &GoroutineLimit{
		num: num,
		ch:  make(chan struct{}, num),
	}
}

//
// Run
// @Description: run a goroutine
//
func (g *GoroutineLimit) Run(f func()) {
	g.ch <- struct{}{}
	go func() {
		f()
		<-g.ch
	}()
}

//func main() {
//	num := 10
//	goroutineLimit := concurrence.NewGoroutineLimit(2)
//	for i := 0; i < num; i++ {
//		concurrence.Wg.Add(1)
//		value := i
//		goFunc := func() {
//			//do something
//			fmt.Println("value:", value)
//			time.Sleep(time.Second)
//			concurrence.Wg.Done()
//		}
//		goroutineLimit.Run(goFunc)
//	}
//	concurrence.Wg.Wait()
//}
