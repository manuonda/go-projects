package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"tutorial/hello/mathutils"
)

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

func printMessage(wg *sync.WaitGroup, message string) {
	defer wg.Done() // Indica que la goroutine ha terminado
	fmt.Println(message)
}

/*
*

	channels
	 <- chan int: solo de se puede leer en el canal
	 chan <- int: solo se puede escribir en el canal
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Trabajador %d procesando trabajo %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Trabajador %d finalizo trabajo %d\n", id, j)
		results <- j * 2
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

	var wg sync.WaitGroup
	wg.Add(1) // Agrega una goroutine al grupo
	go printMessage(&wg, "Hola desde una goroutine!")
	wg.Wait() // Espera a que todas las goroutines terminen
	fmt.Println("Hola desde la funcion prinicipal!")

	print("== Channels ===")
	messages := make(chan string)
	go func() {
		messages <- "Hola desde el canal"
	}()
	msg := <-messages
	fmt.Println("\n", msg)

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		fmt.Println("Resultado : ", <-results)
	}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hola"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Mundo"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recibido ", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recibido ", msg2)
		}
	}

	fmt.Print("=== Testing ===")
	result, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error : ", err)
		//return
	}
	fmt.Println("Resultado : ", result)

	_, err = sqrt(-4)
	if err != nil {
		fmt.Print("Error : ", err)
	}

	var resultTwo = mathutils.Add(1, 2)
	fmt.Printf("\n1 + 2  = %d\n", resultTwo)
}

type ErrorDetail struct {
	Arg  float64
	Prob string
}

/**
* Implementa la interfaz error
 */
func (e *ErrorDetail) Error() string {
	return fmt.Sprintf("%f - %s", e.Arg, e.Prob)
}

/**
* Funcion que retorna la raiz cuadrada de un numero
* tiene dos valores de retorno
* float64: resultado
* error: error es una interfaz
 */

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, &ErrorDetail{x, "No puede ser numero negativo"}
	}
	return x * x, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil
}
