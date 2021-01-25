package main

import "fmt"

/*
An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.

funtion to rotate:

func rotate(A []int, K int) []int

that, given an array A consisting of N integers and an integer K, returns the array A rotated K times.

For example, given

    A = [] int{3, 8, 9, 7, 6}
    K = 3
the function should return [9, 7, 6, 3, 8]. Three rotations were made:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
For another example, given

    A = [] int{0, 0, 0}
    K = 1
the function should return [0, 0, 0]

Given

    A = [] int {1, 2, 3, 4}
    K = 4
the function should return [1, 2, 3, 4]

Assume that:

N and K are integers within the range [0..100];
each element of array A is an integer within the range [−1,000..1,000].
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.

*/

func main() {
	arr := []int{3, 8, 9, 7, 6}
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
