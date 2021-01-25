package main

import "fmt"

func main() {
	// odd list
	odd := []int{10, 20, 30, 40, 50}

	// evn list
	evn := []int{10, 20, 30, 40}
	fmt.Println("Before list:", odd)
	revers(odd)
	fmt.Println("odd list:", odd)
	fmt.Println("Before list:", evn)
	revers(evn)
	fmt.Println("even list:", evn)

}

func revers(arr []int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-(i+1)] = arr[l-(i+1)], arr[i]
	}
}
