package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/manuonda/go-projects/src/cli-todo"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add new todo")
	flag.Parse()

	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
	default:
		fmt.Println(os.Stdout, "invalid command")
		os.Exit(0)
	}

}
