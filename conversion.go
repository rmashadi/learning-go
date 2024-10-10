package main

import "fmt"

// di golang kadang membutuhkan konversi tipe data dari satu tipe ke tipe yang lain
// Misal kita ingin mengonversi tipe data int32 ke int63, dan yang lainnya.

func main() {
	// #1

	//var nilai32 int32 = 32767
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// #2
	var name = "mashans"
	var e = name[0] // dalam bentuk byte
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
