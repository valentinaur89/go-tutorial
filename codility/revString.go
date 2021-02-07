package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]
	var revStr string
	for _, v := range str {
		revStr = string(v) + revStr
	}
	fmt.Println(revStr)
}
