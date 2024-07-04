package main

import (
	"fmt"
	"slices"
)

func DeleteDuplicates(arr []int) (result []int) {
	for _, value := range arr {
		if !slices.Contains(result, value) {
			result = append(result, value)
		}
	}
	return
}

func main() {
	arr := []int{2, 2, 3, 3, 5, 6, -1, 9, 6, 10, -1}
	arr = DeleteDuplicates(arr)
	fmt.Println(arr)
}
