package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Instructor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type WorkShop struct {
	Title       string       `json:"title"`
	Instructors []Instructor `json:"instructors"`
}

func main() {
	fileName := flag.String("file", "", "Input file name to read")
	flag.Parse()
	if *fileName == "" {
		fmt.Println(os.Stderr, "Please specify file name", flag.Args())
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(os.Stderr, "Could not open file", flag.Args())
		os.Exit(1)
	}

	var jsonData WorkShop

	err = json.NewDecoder(file).Decode(&jsonData)

	if err != nil {
		fmt.Println(os.Stderr, "Could not unmarshall decode", err.Error())
		os.Exit(1)
	}

	fmt.Println("WorkShop : %s", jsonData.Title)
	fmt.Println("Instructors : ")
	for _, instructor := range jsonData.Instructors {
		fmt.Println(" %s (%s)", instructor.Name, instructor.Email)
	}

}
