package main

import (
	"fmt"
	"strings"
)

func IsVowel(symbol string) bool {
	symbol = strings.ToLower(symbol)
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	for _, vowel := range vowels {
		if symbol == vowel {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(IsVowel("a"))
	fmt.Println(IsVowel("U"))
	fmt.Println(IsVowel("b"))
	fmt.Println(IsVowel("1"))
}
