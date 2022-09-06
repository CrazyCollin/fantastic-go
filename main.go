package main

import (
	"fantastic-go/concurrence"
	"fmt"
	"time"
)

func main() {
	//fmt.Printf("%d\n", interview.TestDefer())

	//interview.TestReflect()
	//interview.TestSet()
	//interview.TestStringJoin()

	num := 10
	goroutineLimit := concurrence.NewGoroutineLimit(2)
	for i := 0; i < num; i++ {
		concurrence.Wg.Add(1)
		value := i
		goFunc := func() {
			//do something
			fmt.Println("value:", value)
			time.Sleep(time.Second)
			concurrence.Wg.Done()
		}
		goroutineLimit.Run(goFunc)
	}
	concurrence.Wg.Wait()
}
