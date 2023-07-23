package util

import (
	"flag"
	"fmt"
	"os"

	"github.com/manuonda/go-projects/exercises_go/example7/models"
	"gopkg.in/yaml.v3"
)

func Transform(fileName string) (models.WorkShop, error) {

	var jsonData models.WorkShop

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
