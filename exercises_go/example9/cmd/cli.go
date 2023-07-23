package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	util "github.com/manuonda/go-projects/exercises_go/example7/utils"
)

func main() {
	fileName := flag.String("file", "", "Input file name to read")
	silence := flag.Bool("silence", false, "Silence the output")
	output := flag.String("output", "", "Output file name to write")

	flag.Parse()
	if *fileName == "" {
		fmt.Println("Please specify file name", flag.Args())
		os.Exit(1)
	}

	yamlData, err := util.Transform(*fileName)

	if err != nil {
		fmt.Println("Error transforming data to yaml", err.Error())
		os.Exit(1)
	}
	fmt.Println("paoso 1")

	fmt.Println(*silence)
	if !*silence {
		fmt.Println(yamlData)
	}

	if *output != "" {
		err := ioutil.WriteFile(*output, []byte(fmt.Sprint(yamlData)), 0644)
		if err != nil {
			fmt.Println("Error writing file", err.Error())
			os.Exit(1)
		}
	}

}
