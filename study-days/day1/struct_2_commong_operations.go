package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method on the Person struct
func (p *Person) sayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old .\n", p.Name, p.Age)
}

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	alice := Person{Name: "Alice", Age: 30}
	fmt.Println("Struct Instance : ", alice)

	//Accessing and modifying fields
	fmt.Println("Name before: ", alice.Name)

}
