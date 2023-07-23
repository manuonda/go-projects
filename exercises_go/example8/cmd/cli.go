package main

import (
	"flag"
	"fmt"
	"os"

	util "github.com/manuonda/go-projects/exercises_go/example7/utils"
)

func main() {
	fileName := flag.String("file", "", "Input file name to read")
	flag.Parse()
	if *fileName == "" {
		fmt.Println("Please specify file name", flag.Args())
		os.Exit(1)
	}

	jsonData, err := util.Transform(*fileName)

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
