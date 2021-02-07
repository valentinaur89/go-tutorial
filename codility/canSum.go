package main

import "fmt"

/*
*Given an array of numbers and a separate number,
*how would you determine the first combination of 2 numbers
*in that array that would total this single other number?
 */
func main() {
	a := []int{100, 100, 600, 100}
	target := 900
	fmt.Println(canSum(target, a))
}

func canSum(target int, nums []int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{nums[idx], v}
		}
		m[v] = k
	}
	return nil
}
