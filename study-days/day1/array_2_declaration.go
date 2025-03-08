package main

import "fmt"

func main() {

	// 1. Declaration without initialization
	var array1 [5]int
	fmt.Println("Array declared without initialization : ", array1)

	// 2. Declaration with initialization
	var array2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array declared : ", array2)

	//3. short declaration
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array with short hand declaration: ", arr3)

	// 4 . Initializing with specific elements
	arr4 := [5]int{1: 10, 3: 30}
	fmt.Println("Array initialing with specific elements :", arr4)

	// 5. Using the ... Operator to Inter Length
	arr5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Array using .... operator to infer length : ", arr5)

	// 6 . Array of Arrays (Multi-dimensional Arrays)
	var multiArray [2][3]int
	multiArray[0] = [3]int{1, 2, 3}
	multiArray[1] = [3]int{4, 5, 6}

	// 7 . Initializing Multi-dimensional Arrays
	initializezMultiArray := [2][3]int{{7, 8, 9}, {10, 11, 12}}
	fmt.Println("Initialized Multi-Dimensional Array :", initializezMultiArray)

	// 8 . Arrays with Pointers
	var pointerArray [3]*int
	num1, num2, num3 := 13, 14, 15
	pointerArray[0] = &num1
	pointerArray[1] = &num2
	pointerArray[2] = &num3
	fmt.Println("Array with pointers")
	for i := 0; i < len(pointerArray); i++ {
		fmt.Printf("pointerArray[%d] = %d\n", i, *pointerArray[i])
	}

	// 9 . Arrays of Structs
	type Person struct {
		Name string
		Age  int
	}
	var structArray [2]Person
	structArray[0] = Person{"Alice", 23}
	structArray[1] = Person{"Bob", 13}
	fmt.Println("Array of structs: ", structArray)
}
