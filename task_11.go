package main

import "fmt"

func ConvertToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	fmt.Println(ConvertToFahrenheit(30.0))
}
