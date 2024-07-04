package main

import "fmt"

func Max(x, y, z int) int {
	if x > y && x > z {
		return x
	} else if y > x && y > z {
		return y
	} else {
		return z
	}
}

func main() {
	fmt.Println(Max(99, -13, 55))
}
