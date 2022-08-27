package main

import (
	"fmt"
	"math"
)

func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Prime(3))
	fmt.Println(Prime(4))
	fmt.Println(Prime(49))
	fmt.Println(Prime(77))
	fmt.Println(Prime(103))
	fmt.Println(Prime(361))
}
