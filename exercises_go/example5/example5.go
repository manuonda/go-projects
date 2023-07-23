package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Instructor struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type WorkShop struct {
	Title       string       `yaml:"title"`
	Instructors []Instructor `yaml:"instructors"`
}

func main() {
	fileName := flag.String("file", "", "Input file name to read")
	flag.Parse()
	if *fileName == "" {
		fmt.Println("Please specify file name", flag.Args())
		os.Exit(1)
	}

	file, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Could not Read File", flag.Args())
		os.Exit(1)
	}

	var jsonData WorkShop

	err = yaml.Unmarshal(file, &jsonData)

	if err != nil {
		fmt.Println("Could Not Unmarshall Decode", err.Error())
		os.Exit(1)
	}

	fmt.Println("WorkShop : %s", jsonData.Title)
	fmt.Println("Instructors : ")
	for _, instructor := range jsonData.Instructors {
		fmt.Println(" %s (%s)", instructor.Name, instructor.Email)
	}

}
