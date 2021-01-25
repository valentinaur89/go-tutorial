package main

import "fmt"

func main() {
	arr := []int{100, 74, 20, 10, 40, 80, 3, 2, 41}
	fmt.Println("Before v1:", arr)
	bubbleV1(arr)
	fmt.Println("After v1:", arr)
	arr = []int{100, 74, 20, 10, 40, 80, 3, 2, 41}
	fmt.Println("Before v2:", arr)
	bubbleV2(arr)
	fmt.Println("After v2:", arr)
}

// version 1
func bubbleV1(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

//version 2 refined : incase of sorted array we are skipping iteration
func bubbleV2(arr []int) {
	var isSorted bool
	for i := 0; i < len(arr); i++ {
		isSorted = true
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				isSorted = false
			}

		}
		if isSorted {
			return
		}
	}
}
