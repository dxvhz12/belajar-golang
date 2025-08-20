package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var NoKTPUdin NoKTP = "1286386132"
	var MarriedUdin Married = true

	fmt.Println(NoKTPUdin)
	fmt.Println(MarriedUdin)

}