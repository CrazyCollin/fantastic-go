package interview

import "fmt"

func TestDefer() int {
	i := 0
	defer func() {
		fmt.Printf("i=%d\n", i)
	}()
	defer func() {
		i++
		fmt.Printf("i=%d\n", i)
	}()
	return i
}

// Output:
// i=1
// i=1
// 0
