package main

import "fmt"

func main() {
	var x int = 5

	fmt.Println(x)

	var y func() = func() { // cetak
		fmt.Println("hello world")
	}

	y()
}

// func cetak() {
// 	fmt.Println("hello world")
// }