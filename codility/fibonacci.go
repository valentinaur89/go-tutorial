package main

import (
	"fmt"
	"log"
	"math/big"
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

//Using recurrsion : good until 50 fib
func fibrecurrsion(num int) int {
	if num <= 1 {
		return num
	}
	return fibrecurrsion(num-1) + fibrecurrsion(num-2)
}

//Using Iteration & dynamic programing
func fibiteration(num int) *big.Int {
	fib := make([]*big.Int, num+1, num+2)
	if num < 2 {
		fib = fib[0:2]
	}
	fib[0] = big.NewInt(0)
	fib[1] = big.NewInt(1)
	for i := 2; i <= num; i++ {
		fib[i] = (fib[i-1]).Add(fib[i-2])
	}
	return fib[num]
}
