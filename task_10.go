package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	r := Rectangle{width: 14, height: 56}
	fmt.Println(r.Area())
}
