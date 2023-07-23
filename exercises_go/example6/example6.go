package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
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
func tarnsform(fileName string) (WorkShop, error) {

	var jsonData WorkShop

	file, error := os.ReadFile(fileName)
	if error != nil {
		fmt.Println("Could not Read File", flag.Args())
		os.Exit(1)
	}

	error = yaml.Unmarshal(file, &jsonData)

	if error != nil {
		fmt.Println("Could Not Unmarshall Decode", error.Error())
		os.Exit(1)
	}

	return jsonData, nil
}
