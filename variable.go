package main

import "fmt"

func main() {
	// var name string

	// name = "Mas Hans"
	// fmt.Println(name)

	// name = "mashan"
	// fmt.Println(name)

	name := "Mas Hans" // := (pertama kali saja) untuk menggantikan var name =
	fmt.Println(name)

	name = "mashan"
	fmt.Println(name)

	// Deklarasi Multiple Variable
	// * Di Go-lang kita bisa membuat variable secara banyak sekaligus
	// * Code yang dibuat akan lebih bagus dan mudah dibaca

	var (
		firstName = "mas"
		lastName  = "hans"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
