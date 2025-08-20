package main

import "fmt"

func main() {
	// konversi integer
	var ini32 int32 = 129
	var ini64 int64 = int64(ini32)
	var ini8 int8 = int8(ini32)

	fmt.Println(ini32)
	fmt.Println(ini64)
	fmt.Println(ini8)

	// konversi string
	var name = "Udin"
	var g byte = name[0]
	var gString string = string(g)

	fmt.Println(name)
	fmt.Println(gString)
}