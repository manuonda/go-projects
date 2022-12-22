package main

import "fmt"

func main() {
	fmt.Print("Kindly enter your data in following order:")
	var firstName, lastName, occupation, country string
	var phoneNo int64
	fmt.Print("firstname, lastname, occupation, country, phoneNo")
	totalInputs, errReading := fmt.Scanln(&firstName, &lastName, &country, &occupation, &phoneNo)

	if errReading != nil {
		fmt.Printf("Your input could not be read%v", errReading)
	}

	// fullname
	fullname := fmt.Sprintf("%v %v ", firstName, lastName)
	fmt.Printf("totalinputs : %v ", totalInputs)
	fmt.Printf("Full name %v", fullname)
}
