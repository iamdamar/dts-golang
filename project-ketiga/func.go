package main

import "fmt"

func main() {
	cetakNama("budi", "ani", "joko", "asdf")
}

func cetak() {
	fmt.Println("hello world")
}

func cetakNama(name ...string) {
	fmt.Println("hello", name)
}

func sum(a, b int) int {
	return a + b
}