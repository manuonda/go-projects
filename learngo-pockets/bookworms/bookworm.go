package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bookworm struct {
	Name  string  `json:"name"`
	Books []Books `json:"books"`
}

type Books struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func LoadBookworms(filePath string) ([]Bookworm, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm

	// decode de file and store the content in the value bookworms
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

func Imprimir() {
	fmt.Print("hola mundo")
}
