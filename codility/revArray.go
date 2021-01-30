package main

import "fmt"

func main() {

	i := make([]interface{}, 0, 0)
	i = append(i, 10, "a", 30, 40)
	fmt.Printf("Before ::%v ", i)
	revers(i)
	fmt.Printf("After::%v\n", i)
	i = i[:0]
	i = append(i, 1, 2, 3, "c")
	fmt.Printf("Before ::%v After::%v\n", i, reverseRec(i))
}

// with loop
func revers(arr []interface{}) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-(i+1)] = arr[l-(i+1)], arr[i]
	}

}

// with recurs
func reverseRec(arr []interface{}) []interface{} {
	if len(arr) == 0 {
		return arr
	}

	return append(reverseRec(arr[1:]), arr[0])
}
