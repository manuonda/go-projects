package main

import "fmt"

func modify(s []int) {
	if len(s) > 0 {
		s[0] = 999 //Modify the first element
	}
	fmt.Println("Inside modifySlice, after modification", s)
}

func main() {

	//create a slice
	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice before modification : ", originalSlice)

	//Pass the slice to the modifySlice function
	modify(originalSlice)

	//Print the slice again in the main function
	fmt.Println("Original slice after modification : ", originalSlice)
}
