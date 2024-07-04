package main

import "fmt"

func MinMax(arr []int) (int, int) {
	minel := arr[0]
	maxel := arr[0]
	for _, x := range arr {
		if x < minel {
			minel = x
		} else if x > maxel {
			maxel = x
		}
	}
	return minel, maxel
}

func main() {
	arr := []int{-3, 88, 102, 1, 0, -26, 33, 999}
	fmt.Println(MinMax(arr))
}
