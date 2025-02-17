package main

import "fmt"

type Address struct {
	City, State string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Rectangle struct {
	Widht  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	r.Widht = 20 // No modifica el valor de la estructura
	return 2 * (r.Height + r.Widht)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

// Toma el valor original de la estructura
func (r *Rectangle) Scale(factor float64) {
	r.Widht *= factor
	r.Height *= factor
}

// Interfaz Shape
type Shape interface {
	Area() float64
}

type Engine struct {
	HorsePower int
}

func (e Engine) Start() {
	fmt.Println("Engine started with ", e.HorsePower, "HorsePower")
}

type Car struct {
	Make  string
	Model string
	Engine
}

func doSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown type")
	}

}

func main() {
	fmt.Println("Hell world")
	var float32 float32 = 3.14
	var float64 float64 = 3.14

	fmt.Printf("float32: %.10f\n", float32) // Imprime 10 decimales
	fmt.Printf("float64: %.17f\n", float64) // Imprime 17 decimales

	sayHello()

	c := counter()
	fmt.Println(c()) //imprime 1
	fmt.Println(c()) //imprime 2

	x := 10
	cambiarValor(&x)
	fmt.Println(x)

	person := Person{
		Name: "Juan",
		Age:  25,
		Address: Address{
			City:  "Lima",
			State: "NY",
		},
	}
	fmt.Println(person)

	r := Rectangle{Widht: 10, Height: 5}
	fmt.Println("Area rectangle ", r.Area())

	//r.Scale(2.5)
	//fmt.Println("Scale ", r)
	circle := Circle{Radius: 5}
	printArea(r)
	printArea(circle)

	doSomething(10)
	doSomething("Hello")
	doSomething(3.14)

	car := Car{
		Make:  "Toyota",
		Model: "Corolla",
		Engine: Engine{
			HorsePower: 132,
		},
	}
	car.Start()
}

func sayHello() {
	fmt.Println("Hello")
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func cambiarValor(p *int) {
	*p = 42
}
