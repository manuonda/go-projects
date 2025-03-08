package main

import "fmt"

func modifyArray(arr [5]int, index int, value int) [5]int {
	arr[index] = value
	return arr
}

func main() {
	//initialize array
	originalArray := [5]int{1, 2, 3, 4, 5}

	//Print the length of the array
	fmt.Printf("Length of the array is  : ", len(originalArray))

	//Copy the array and print the copied array
	var copiedArray = originalArray //This creates a copy
	fmt.Println("Copied array : ", copiedArray)

	//Modify the third element (index 2) of the original
	modifiedArray := modifyArray(originalArray, 2, 10)
	fmt.Println("Modified Array: ", modifiedArray)

	//Compare the original and modified arras
	isEqual := originalArray == modifiedArray
	fmt.Println("Are original and modified arrays equals?", isEqual)

	isCopiedEqual := originalArray == copiedArray
	fmt.Println("Are original and copied arrays equals: ", isCopiedEqual)
	fmt.Println("original array : ", originalArray)
	fmt.Println("copiedArray : ", copiedArray)
}
