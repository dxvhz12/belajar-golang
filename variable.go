package main

import "fmt"

func main() {

	// versi 1
	var name string
	var number int

	name = "Udin"
	fmt.Println(name)

	name = "Udin Kurniawan"
	fmt.Println(name)

	number = 1
	fmt.Println(number)

	// versi 2
	var nickname = "Jalaludin"
	fmt.Println(nickname)

	var kodePos = 17716
	fmt.Println(kodePos)

	// versi 3
	myName := "Udin Jalaludin"
	fmt.Println(myName) 

	myName = "Udin Bahrudin"
	fmt.Println(myName) 

	// versi 4 (multiple var)
	var (
		myFriend = "Udindang"
		age = 11
	)

	fmt.Println(myFriend)
	fmt.Println(age)
}