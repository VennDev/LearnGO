package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func mul(x int, y int) int {
	return x * y
}

// Function with multiple return values
func add(x int, y int) (int, int) {
	return x + y, x - y
}

// Function with named return values
func div(x int, y int) (xl int, yl int, isSquare bool) {
	xl = x / y
	yl = x % y
	isSquare = x == y
	return xl, yl, isSquare
}

func main() {
	hello()
	fmt.Println(mul(10, 20))
	fmt.Println(add(10, 20))
	fmt.Println(div(10, 20))
	fmt.Println(div(10, 10))
}
