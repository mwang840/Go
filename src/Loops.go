package main

import (
	"fmt"
	"time"
)

func main() {
	//Looping via a single condition
	i := 0
	for i <= 4 {
		fmt.Println(i)
		i += 1
	}

	//Looping with an initial single condition
	for j := 5; j < 10; j++ {
		fmt.Println(j)
	}
	//Loops with breaks
	for {
		fmt.Println("Exiting out loop!")
		break
	}

	//Loops with continue upon condition
	for k := 10; k <= 15; k++ {
		if k%3 == 0 {
			continue
		}
		fmt.Println(k)
	}

	//Switch casing
	l := 6
	fmt.Print("Say ", l, " as ")
	switch l {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	}
	//Switch casing with time
	//In this case it checks the day of the week (open mouth)
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//Switch cases with types
	myType := func(i interface{}) {
		switch s := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("I don't know what type I am %T\n", s)
		}
	}
	myType(8.1)
	myType(false)
	myType("Ayo")
	myType(10)
}
