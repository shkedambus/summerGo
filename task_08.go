package main

import "fmt"

func ReverseString(s string) (reversed string) {
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return
}

func main() {
	fmt.Println(ReverseString("Hello, world!"))
}
