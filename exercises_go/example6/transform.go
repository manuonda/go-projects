package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

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
