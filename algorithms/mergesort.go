package main

import (
	"fmt"
)

func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {

		if len(left) == 0 {

			return append(res, right...)
		}
		if len(right) == 0 {

			return append(res, left...)
		}
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	return res
}
func main() {
	arr := []int{40, 80, 10, 5, 35, 70, 50}
	fmt.Println("Before sort ::", arr)
	fmt.Println("After sort :: ", mergesort(arr))
}
