package main

/**
Lo interesante en Go es que no necesitas declarar
explícitamente que una estructura implementa una interfaz.
Si una estructura tiene todos los métodos requeridos
por una interfaz (como Dog tiene el método Speak()
requerido por la interfaz Speaker), entonces automáticamente
implementa esa interfaz.

Esta es una característica de diseño que hace que Go sea muy flexible y
promueve un acoplamiento débil entre componentes.

**/

import "fmt"

type Speaker interface {
	Speak() string
	Read() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + "say Woof!"
}

func (d Dog) Read() string {
	return d.Speak() + " speak Dog!"
}

// 2. Embbeding Interfaces
type Animal interface {
	Speaker
	Move() string
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return b.Name + "says Tweet"
}

func (b Bird) Read() string {
	return b.Name + " bird Speak"
}

func (b Bird) Move() string {
	return b.Name + " flies"
}

// 3. Empty Interface
func PrintAnything(v interface{}) {
	fmt.Println(v)
}

// 4. Interface with multiple Methods
type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Start() string {
	return c.Model + " car stopped"
}

func (c Car) Stop() string {
	return c.Model + " car stopped "
}

// 5  . Functional Interface(SIngle Method Interface)
//Similar to Speaker interface

// 6. Interface as a Constraint(Generics)
func Describe[T Speaker](t T) {

	fmt.Println(t.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	bird := Bird{Name: "Tweety"}
	car := Car{Model: "Tesla"}

	fmt.Println("Animal Dog --")
	fmt.Println(dog.Speak())
	fmt.Println(dog.Read())
	fmt.Println(bird.Speak())
	fmt.Println(bird.Move())

	PrintAnything("An empty interface can hold anything")
	PrintAnything(42)

	fmt.Println(car.Start())
	fmt.Println(car.Stop())

	Describe(dog)
	Describe(bird)

}
