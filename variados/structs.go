package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{}
	p.name = name
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"BOB", 43})
	fmt.Println(person{name: "Alice", age: 43})
	fmt.Println(person{name: "Fred", age: 23})
	fmt.Println(&person{name: "David agarcia", age: 37})
	fmt.Println(newPerson("Andres"))

}
