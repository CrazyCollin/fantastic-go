package concurrence

import (
	"fmt"
)

//分别打印cat，fish，dog，分别用一个协程解决

var CatChan = make(chan struct{})
var FishChan = make(chan struct{})
var DogChan = make(chan struct{})

func Cat() {
	<-DogChan
	fmt.Printf("cat\n")
	CatChan <- struct{}{}
}

func Fish() {
	<-CatChan
	fmt.Printf("fish\n")
	FishChan <- struct{}{}
}

func Dog() {
	<-FishChan
	fmt.Printf("dog\n")
	DogChan <- struct{}{}
}

//func main() {
//	for i := 0; i < 100; i++ {
//		go concurrence.Cat()
//		go concurrence.Fish()
//		go concurrence.Dog()
//	}
//	concurrence.DogChan <- struct{}{}
//	time.Sleep(2 * time.Second)
//}
