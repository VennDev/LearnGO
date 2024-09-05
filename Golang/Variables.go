package main

import (
	"fmt"
	"math"
	"math/bits"
	"unsafe"
)

func main() {

	// Addition of two numbers
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println(c)

	// Addition of two strings
	var str1 string = "Hello"
	var str2 string = "World"
	var str3 string = str1 + str2
	fmt.Println(str3)

	// Addition of two booleans
	var bool1 bool = true
	var bool2 bool = false
	var bool3 bool = bool1 && bool2
	fmt.Println(bool3)

	// Autoincrement
	var i int = 10
	i++
	fmt.Println(i)

	// Autodecrement
	var j int = 10
	j--
	fmt.Println(j)

	// Auto detect type
	var k = 10
	fmt.Println(k)

	// Auto detect type
	l := 10
	fmt.Println(l)

	// Constant
	const pi float32 = 3.14
	fmt.Println(pi)

	// Bits and Memory
	var aa int = 10
	fmt.Printf("Type: %T, Size: %d, Value: %v\n", aa, unsafe.Sizeof(aa), aa)

	// Bits
	fmt.Println(bits.OnesCount8(math.MaxUint8))
}
