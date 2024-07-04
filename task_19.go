package main

import "fmt"

func Delete(i int, arr []int) (result []int) {
	for index, value := range arr {
		if i != index {
			result = append(result, value)
		}
	}
	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = Delete(2, arr)
	fmt.Println(arr)
	arr = Delete(6, arr)
	fmt.Println(arr)
}
