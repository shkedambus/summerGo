package main

import "fmt"

func Count(x int, arr []int) int {
	cnt := 0
	for _, value := range arr {
		if x == value {
			cnt += 1
		}
	}
	return cnt
}

func main() {
	arr := []int{-3, 5, 3, 29, 5, -1, 3, 6, 8, 3, 5, 8, -19, 3}
	fmt.Println(Count(3, arr))
}
