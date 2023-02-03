package main

import "fmt"

func main() {
	//initializes an empty string array with a given size
	str := make([]string, 4)
	fmt.Println("Empty set:", str)

	str[0] = "T"
	str[1] = "h"
	str[2] = "a"
	str[3] = "t"
	//Prints out the filled string, the string at the first element and the length of the first array
	fmt.Println("Set: ", str)
	fmt.Println("Get: ", str[1])
	fmt.Println("Length: ", len(str))

	//Sets up a new array with the same length
	word := make([]string, len(str))
	copy(word, str)
	fmt.Println("cpy:", word)

}
