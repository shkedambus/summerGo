package main

import "fmt"

func MultiplicationTable(n int) int {
	for i := 1; i <= 9; i++ {
		fmt.Printf("%v * %v = %v\n", n, i, n*i)
	}
	return n
}

func main() {
	MultiplicationTable(34)
}
