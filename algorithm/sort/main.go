package main

import "fmt"

func main() {
	arr := []int{5, 2, 3, 1}
	//quickSort(arr, 0, len(arr)-1)
	//mergeSort(arr, 0, len(arr)-1)
	heapSort(arr)
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
}
