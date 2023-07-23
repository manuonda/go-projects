package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/*
*
Example 2 - Create a CLI that takes in a file input,

	reads in the file, then outputs the content

*
*/
func main() {
	fileName := flag.String("file", "", "File name to read from")
	flag.Parse()
	if *fileName == "" {
		fmt.Println(os.Stderr, "Error file name is required ", flag.Args())
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(os.Stderr, "Error opening file: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error reading file:", err.Error())
		os.Exit(1)
	}

}
