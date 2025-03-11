package main

import "fmt"

func main() {
	//Initialize a map
	fruits := map[string]int{"apple": 5, "banana": 3}

	//add an element
	fruits["orange"] = 2
	for key, value := range fruits {
		fmt.Printf("%s: %d\n", key, value)
	}
}
