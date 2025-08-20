package main

import "fmt"

func main() {
	// versi 1
	const admin string = "udin"
	const friend = "andi"
	const age = 14

	fmt.Println(admin)
	fmt.Println(friend)
	fmt.Println(age)

	// versi 2 (multiple const)
	const (
		subAdmin string = "jalu"
		number = 2
	)

	fmt.Println(subAdmin)
	fmt.Println(number)
}