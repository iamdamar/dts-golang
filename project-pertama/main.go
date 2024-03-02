package main

import "fmt"

func main() {
	fmt.Println("Hello world", "Golang", 123) // ini komentar
	fmt.Println("Hello world")

	var name string = "Budi" // buat variable cara 1
	fmt.Println(name)

	var address = "Jakarta" // buat variable cara 2
	fmt.Println(address)

	product := "iphone" // buat variable cara 3
	fmt.Println(product)

	var total int // buat variable cara 4
	total = 10
	fmt.Println(total)

	tidak := "vava" // buat variable tanpa melakukan print atau operation
	_ = tidak

	fmt.Printf("%T\n", product) // melihat tipe data menggunakan printf
}
