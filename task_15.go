package main

import "fmt"

func Average(arr []int) float64 {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Average(arr))
}
