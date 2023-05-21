package main

import (
	"fmt"
	"os"
)

func main() {

	bookworms, err := LoadBookworms("testdata/bookworm.json")
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "failed to load boolworms : %s", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the books in common : ")
	displayBooks(commonBooks)

	//sortBooks : = sortsortBooks()

}

func Imprimi2() {
	fmt.Println("imprimi 2")
}
