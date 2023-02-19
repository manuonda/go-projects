package main

import (
	"fmt"

	"github.com/manuonda/go-projects/tutorials/interface/users"
)

func main() {
	var frank users.Cashier = users.NewEmployee("Frank")
	var robert users.Admin = users.NewEmployee("David")

	total := frank.CalcTotal(90, 23, 32, 43)
	fmt.Println(total)

	robert.DeactivateEmplee(frank)
	fmt.Println(frank.CalcTotal(23, 11, 12))

}
