package main

import "fmt"

// Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
// Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
// Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya.

func main() {
	// #1
	// const firstName string = "Riska"
	// const lastName = "Handika"

	// error

	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"

	// #2
	//Deklarasi Multiple Constant
	// * Sama seperti variable. di Go-lang juga kita bisa membuat constant secara banyak sekaligus

	const (
		firstName string = "Riska"
		lastName         = "Handika"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
