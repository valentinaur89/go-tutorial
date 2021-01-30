package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	fmt.Print("Recurrsive:")
	for i := 0; i < num; i++ {
		fmt.Print(fibrecurrsion(i), " ")
	}
	fmt.Println()
	fmt.Print("Iterative:")
	for i := 0; i < num; i++ {
		fmt.Print(fibiteration(i), " ")
	}
	fmt.Println()

}

func fibrecurrsion(num int) int {
	if num <= 1 {
		return num
	}
	return fibrecurrsion(num-1) + fibrecurrsion(num-2)
}

func fibiteration(num int) int {
	fib := make([]int, num+1, num+2)
	if num < 2 {
		fib = fib[0:2]
	}
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= num; i++ {

		fib[i] = fib[i-1] + fib[i-2]

	}
	return fib[num]
}
