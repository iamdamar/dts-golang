package main

import "fmt"

func main() {
	var age = 10

	if age > 17 {
		fmt.Println("punya KTP")
	} else {
		fmt.Println("belum punya KTP")
	}

	if age > 17 {
		fmt.Println("punya KTP")
	} else {
		fmt.Println("punya kartu pelajar")
	}

	switch age { // switch mirip seperti if age == 10
	case 10:
		fmt.Println("ini 10")
	default:
		fmt.Println("ini bukan 10")
	}
}