package main

import "fmt"

func Contains(x int, arr []int) bool {
	for _, element := range arr {
		if element == x {
			return true
		}
	}
	return false
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

	fmt.Println(Contains(x, arr))
}
