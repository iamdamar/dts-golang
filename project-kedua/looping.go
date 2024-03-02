package main

import "fmt"

func main() {
	// cara 1
	for index := 1; index <= 5; index++ { // ini adalah temporary declaration, dimana index hanya bisa dipakai dalam satu for ini
		fmt.Println(index)
	}

	fmt.Println("end")

	// cara 2
	index := 1 // ini adalah permanent declaration, dimana index bisa dipakai oleh for yang lain juga
	for ; index <= 5; index++ {
		fmt.Println(index)
	}

	fmt.Println("end")

	// cara 3
	index = 1
	for index <= 5 {
		fmt.Println(index)
		index++
	}

	fmt.Println("end")

	// cara 4
	index = 1
	for {
		if index <= 5 {
			fmt.Println(index)
			index++
		} else {
			break
		}
	}

	fmt.Println("end")
}