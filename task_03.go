package main

import "fmt"

func OddOrEven(x int) string {
	if x%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	fmt.Println(OddOrEven(16))
	fmt.Println(OddOrEven(7))
}
