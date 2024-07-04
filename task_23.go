package main

import "fmt"

func Fibonacci(n int) (sequence []int) {
	sequence = append(sequence, 1, 1)
	if n > 2 {
		for i := 2; i < n; i++ {
			sequence = append(sequence, sequence[i-1]+sequence[i-2])
		}
	}
	return
}

func main() {
	fmt.Println(Fibonacci(20))
}
