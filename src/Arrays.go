package main

import "fmt"

func main() {
	var arr [3]int
	fmt.Println("Empty: ", arr)
	arr[1] = 1
	fmt.Println("Set: ", arr)
	fmt.Println("Get: ", arr[1])
	fmt.Println("Length of array: ", len(arr))

	arr2 := [6]int{0, 2, 4, 6, 8, 10}
	fmt.Println("Inside the array: ", arr2)
}
