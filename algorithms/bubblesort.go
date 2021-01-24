package main

import "fmt"

func main() {
	arr := []int{20, 10, 40, 80, 3, 2, 41}
	fmt.Println("Before:", arr)
	bubble(arr)
	fmt.Println("After:", arr)
}

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
