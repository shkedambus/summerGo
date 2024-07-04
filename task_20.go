package main

import "fmt"

func LinearSearch(x int, arr []int) int {
	for index, value := range arr {
		if x == value {
			return index
		}
	}
	return -1
}

func main() {
	var x int
	fmt.Println("Input a number:")
	fmt.Scanln(&x)

	var n int
	fmt.Println("Input a length of an array:")
	fmt.Scanln(&n)

	var arr []int
	fmt.Println("Input the array:")
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanln(&temp)
		arr = append(arr, temp)
	}

	fmt.Println(LinearSearch(x, arr))
}
