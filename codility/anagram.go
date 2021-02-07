package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

//if two strings are anagram ex: annu, uann
func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Println(s1, " ", s2)
	//Sorted way
	v := func(s1, s2 string) bool {
		if len(s1) != len(s2) {
			return false
		}
		s1 = sortString(s1)
		s2 = sortString(s2)
		if strings.ToLower(s1) != strings.ToLower(s2) {
			return false
		}
		return true
	}(s1, s2)
	fmt.Println("Is anagram", v)
}

func sortString(s string) string {
	val := strings.Split(s, "")
	sort.Strings(val)
	return strings.Join(val, "")
}
