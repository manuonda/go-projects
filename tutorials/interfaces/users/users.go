package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

// TODO: VOLVER A REALIZAR
type Employee struct {
	User
	active bool
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	deactivate()
}

type Admin interface {
	DeactivateEmplee(c Cashier)
}

func NewEmployee(name string) *Employee {

	employe := &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		active: true,
	}

	return employe
}

// se aplica a la estructura la interfaz
func (e *Employee) CalcTotal(item ...float32) float32 {
	var sum float32

	if !e.active {
		fmt.Println("Cashier deactivate")
		return 0
	}

	for _, i := range item {
		sum += i
	}
	return sum * 1.15
}

func (e *Employee) deactivate() {
	e.active = false
}
func (e *Employee) DeactivateEmplee(c Cashier) {
	c.deactivate()
}
