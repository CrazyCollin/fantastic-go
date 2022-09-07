package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 1}
	quickSort(arr, 0, len(arr)-1)
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
}
