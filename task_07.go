package main

import "fmt"

func AllPrimeLessThanN(n int) (prime []int) {
	for m := 2; m < n; m++ {
		flag := true
		for i := 2; i*i < m; i++ {
			if m%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			prime = append(prime, m)
		}
	}
	return
}

func main() {
	fmt.Println(AllPrimeLessThanN(100))
}
