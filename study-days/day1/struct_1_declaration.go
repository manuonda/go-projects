package main

import "fmt"

func main() {
	//Basic declaration of a struct
	type Person struct {
		Name string
		Age  int
	}
	//Declaration and Instantiating togther
	john := Person{Name: "Juan", Age: 30}
	fmt.Println("Declared and Instantied together: ", john)

	//Using the new keyword
	jane := new(Person)
	jane.Name = "Jane"
	jane.Age = 25
	fmt.Println("Using new keyword : ", jane)

	// Anonymous structs
	anon := struct {
		Country string
		Code    int
	}{
		Country: "USA",
		Code:    1,
	}
	fmt.Println("Anonymous struct : ", anon)

	//Nested structs
	type Address struct {
		City  string
		State string
	}
	type Employee struct {
		PersonalInfo Person
		Address      Address
	}

	emp := Employee{
		PersonalInfo: Person{Name: "David", Age: 39},
		Address:      Address{City: "San Salvador de Jujuy", State: "Jujuy"},
	}
	fmt.Println("Nesteds Structs : ", emp)

	//Embedded (Anonymous )Fields
	type Manager struct {
		Person
		Departament string
	}
	mgr := Manager{
		Person:      Person{Name: "Andres", Age: 38},
		Departament: "Departamento one",
	}
	fmt.Println("Embedded fields: ", mgr)

}
