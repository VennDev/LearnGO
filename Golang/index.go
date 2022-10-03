package main

import (
	"fmt"
	"math"
)

const (
	red = iota
	blue
	green
)

func add(a float64, b int) float64 {
	return a + float64(b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	for _, v := range data {
		fmt.Println(v)
	}
}
