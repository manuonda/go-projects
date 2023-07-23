package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("file", "", "Input file name to read")
	flag.Parse()
	if *fileName == "" {
		fmt.Println("Please specify file name", flag.Args())
		os.Exit(1)
	}

	jsonData, err := tarnsform(*fileName)

	if err != nil {
		fmt.Println("Error transforming data to yaml", err.Error())
		os.Exit(1)
	}

	fmt.Println("WorkShop : %s", jsonData.Title)
	fmt.Println("Instructors : ")
	for _, instructor := range jsonData.Instructors {
		fmt.Println(" %s (%s)", instructor.Name, instructor.Email)
	}

}
