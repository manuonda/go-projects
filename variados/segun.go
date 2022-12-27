package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write : ", i, " as")
	switch i {
	case 1:
		fmt.Println(" one")
	case 2:
		fmt.Println(" two")
	case 3:
		fmt.Println(" three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("It's Weekend")
	default:
		fmt.Print("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Print("It's before noon")
	default:
		fmt.Print("It's a after noon")
	}
}
