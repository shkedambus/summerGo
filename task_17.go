package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("aabbccc"))
	fmt.Println(IsPalindrome("AGAga"))
	fmt.Println(IsPalindrome("World"))
}
