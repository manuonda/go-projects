package main

import "fmt"

func main() {
	var matrix [3][3]int

	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = count
			count++
		}
	}

	//print the matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\n", matrix[i][j])
		}
	}
}
