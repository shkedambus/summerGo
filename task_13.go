package main

import "fmt"

func Len(s string) int {
	l := 0
	for index, _ := range s {
		l = index
	}
	return l + 1
}

func main() {
	fmt.Println(Len("Hello, world!"))
}
