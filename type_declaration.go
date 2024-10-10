package main

import "fmt"

//Type Declarations
// * Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada.
// * Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada. dengan tujuan agar lebih mudah dimengerti

func main() {

	type noKTP string

	var KTPmashan noKTP = "11112222333344445"

	var contoh string = "22222222222222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(KTPmashan)
	fmt.Println(contohKTP)
}
