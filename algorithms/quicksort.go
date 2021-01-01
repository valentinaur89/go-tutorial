package main

import (
	"fmt"
)

/*
* 1. Find pivot by partiontioning
* 2. Perform recursive sort
 */
func quicksort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}
func partition(arr []int, low, high int) int {
	pi := arr[high]
	for low < high {
		for low < high && arr[low] < pi {
			low++
		}
		arr[low], arr[high] = arr[high], arr[low]
		for low < high && arr[high] >= pi {
			high--
		}
		arr[low], arr[high] = arr[high], arr[low]
	}
	return low
}

func main() {
	arr := []int{40, 80, 10, 5, 35, 70, 50}
	fmt.Println("Before sort :: ", arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println("After sort :: ", arr)

}
