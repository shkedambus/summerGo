package main

import "fmt"

func Factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func main() {
	fmt.Println(Factorial(5))
}
