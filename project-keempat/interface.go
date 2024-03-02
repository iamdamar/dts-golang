package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) GetName() string {
	return p.Name
}

type Animal interface {
	Walk()
	GetName()
}

type Dog struct{

}

func (d *Dog) Walk() {

}

func (d *Dog) GetName() {}

func main() {
	var a Person
	var b Animal

	fmt.Println(a)
	fmt.Println(b)
}