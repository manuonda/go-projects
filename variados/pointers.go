package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 30
}

func main() {
	i := 1
	fmt.Println("initial : ", i)

	zeroval(i)
	fmt.Println("zeroval : ", i)

	zeroptr(&i)
	fmt.Println("zeroptr change value : ", i)

	fmt.Println("pointer : ", &i)
}
