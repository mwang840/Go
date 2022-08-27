package main

import (
	"fmt"
)

func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(Prime(3))
	fmt.Println(Prime(4))
}
