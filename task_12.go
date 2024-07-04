package main

import "fmt"

func Countdown(n int) int {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
	return n
}

func main() {
	var n int

	fmt.Println("Input a number:")
	fmt.Scanln(&n)

	Countdown(n)
}
