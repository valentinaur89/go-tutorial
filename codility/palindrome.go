package main

import (
	"fmt"
	"os"
	"strings"
)

//parlindrom example: anna,tattarrattat
//go run palindrome.go  tattarrattat
func main() {

	str := os.Args[1]
	l := len(str)
	var isPalindrome bool
	for i := 0; i < l/2; i++ {
		isPalindrome = false
		if strings.EqualFold(string(str[i]), string(str[(l-1)-i])) {
			isPalindrome = true
		}
		if !isPalindrome {
			break
		}
	}
	fmt.Println(isPalindrome)

}
