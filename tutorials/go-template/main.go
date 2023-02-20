package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

var funcs = template.FuncMap{
	"increment": func(num int) int {
		return num + 1
	},
	"decrement": func(num int) int {
		return num - 1
	},
}

func main() {
	//t, err := template.New("greeting").Parse("Hola mi nombre es {{.Name}} tengo {{.Age}}")

	jhon := &Person{Name: "David", Age: 23, Hobbies: []string{"uno", "dos", "tres"}}
	LoadTemplate("greeting.txt", jhon)
}

func LoadTemplate(fileName string, data interface{}) {
	t, err := template.New(fileName).Funcs(funcs).ParseFiles(fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
