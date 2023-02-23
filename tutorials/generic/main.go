package main

import (
	"fmt"

	"github.com/manuonda/go-projects/tutorials/generic/redisclient"
)

type Numeric interface {
	int | float64 | float32
}

// generic
func sum[K string, T Numeric](key K, a T, b T) T {
	return a + b
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name     string
	Position string
}

func main() {
	fmt.Print(sum("key1", 1, 2))
	pp, err := redisclient.Read[[]Person]("persons")
	if err != nil {
		panic(err)
	}

	ee, err := redisclient.Read[[]Employee]("employees")
	if err != nil {
		panic(err)
	}

	fmt.Println("Persons")
	for _, p := range pp {
		fmt.Printf("%s is %d years old", p.Name, p.Age)
	}

	fmt.Println("Employees")
	for _, e := range ee {
		fmt.Printf("%s is %s years old", e.Name, e.Position)
	}

}
