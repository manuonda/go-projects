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
	redisclient.Read([]Person)("persons")
}
