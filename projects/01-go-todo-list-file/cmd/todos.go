package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

var todos Todos
var todosFile = "todos.json"

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("index out of range")
	}
	return nil
}

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
	todos.saveToFile()
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	todos.saveToFile()
	return nil
}

func (todos *Todos) completed(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	todo := &t[index]
	if !todo.Completed {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime

	} else {
		todo.CompletedAt = nil
	}
	todo.Completed = !todo.Completed
	todos.saveToFile()
	return nil
}

func (todos *Todos) list() {
	tbl := table.New(os.Stdout)
	tbl.SetHeaders("Title", "Completed", "CreatedAt", "CompletedAt")

	t := *todos
	for _, todo := range t {
		completed := "No"
		if todo.Completed {
			completed = "Si"
		}
		completedAt := ""
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format(time.RFC3339)
		}
		tbl.AddRow(todo.Title, completed, todo.CreatedAt.Format(time.RFC3339), completedAt)
	}
	tbl.Render()
}

func (todos *Todos) saveToFile() {
	data, err := json.Marshal(todos)
	if err != nil {
		fmt.Printf("Error converting json: %s\n", err.Error())
		panic(err)
	}
	err = os.WriteFile(todosFile, data, 0777)
	if err != nil {
		fmt.Print("Error write file", err)
		panic(err)
	}
}

func (todos *Todos) readFromFile() {
	data, err := os.ReadFile(todosFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Printf("Error readinf file %s\n", err)
		panic(err)
	}
	err = json.Unmarshal(data, todos)
	if err != nil {
		fmt.Printf("Error Unmarshall %s\n", err)
		panic(err)
	}
}
