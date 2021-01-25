package main

import "fmt"

func main() {
	arr := []int{4, 9, 10, 8, 7}
	fmt.Println("Before Rotation ::", arr)
	fmt.Println("Rotate to 3 ::", rotate(arr, 3))

	arr = []int{1, 2, 3, 4}
	fmt.Println("Before Rotation ::", arr)
	fmt.Println("Rotate to 4 ::", rotate(arr, 4))

	arr = []int{0, 0, 0}
	fmt.Println("Before Rotation ::", arr)
	fmt.Println("Rotate to 1 ::", rotate(arr, 1))

}

func rotate(arr []int, k int) []int {
	l := len(arr)
	if k == l || k <= 0 || l == 0 {
		return arr
	}
	val := l - k%l
	sl1 := arr[:val]
	sl2 := arr[val:]
	arr = append(sl2, sl1...)
	return arr

}
