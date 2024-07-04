package main

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{-1, 3, 8, 0, 79, 13, -39, 6}
	BubbleSort(arr)
	fmt.Println(arr)
}
