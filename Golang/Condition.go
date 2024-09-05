package main

import "fmt"

func main() {

	var number int = 10

	if number > 0 {
		fmt.Println("Number is positive")
	} else if number < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// Switch case
	switch number {
	case 10:
		fmt.Println("Number is 10")
	case 20:
		fmt.Println("Number is 20")
	default:
		fmt.Println("Number is not 10 or 20")
	}

	// Switch case with multiple conditions
	switch number {
	case 10, 20, 30:
		fmt.Println("Number is 10, 20 or 30")
	default:
		fmt.Println("Number is not 10, 20 or 30")
	}

	// Switch case with fallthrough
	switch number {
	case 10:
		fmt.Println("Number is 10")
		fallthrough
	case 20:
		fmt.Println("Number is 20")
	default:
		fmt.Println("Number is not 10 or 20")
	}
}
