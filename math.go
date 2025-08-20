package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 20
	var b = 10	
	var c = a * b

	fmt.Println(c)

	// augmented assignment
	var g = 10
	g += 10 // g = g + 10

	fmt.Println(g)
	
	// unary operator
	g++ // g = g + 1

	fmt.Println(g)
	
	g-- // g = g - 1
	
	fmt.Println(g)

	var negative = -25
	var positive = +25

	fmt.Println(negative)
	fmt.Println(positive)
}