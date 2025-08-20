package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80

	var lulusUjian bool = ujian >= 90
	var lulusAbsensi bool = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi

	fmt.Println(lulus)

	fmt.Println(ujian >= 90 && absensi >= 80)
}