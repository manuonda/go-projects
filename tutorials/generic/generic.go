package main

import "fmt"

type Numeric interface {
	int | float64 | float32
}

// generic
func sum[K string, T Numeric](key K, a T, b T) T {
	return a + b
}

func main() {
	fmt.Print(sum("key1", 1, 2))
}
