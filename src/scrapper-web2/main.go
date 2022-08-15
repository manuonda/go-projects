package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var url = "https://tienda.iberiahogar.com.ar/limpieza/?mpage=2&gclid=Cj0KCQjwrs2XBhDjARIsAHVymmSFVwaW0ybAILv2oTJI5shBFpuMm1vjoeIC-9AER4WzX64aF-4Uh5MaAi8QEALw_wcB"
	response, error := http.Get(url)
	if error != nil {
		log.Fatal(error)
	}

	// fmt.Print(response)
	status, err := strconv.ParseInt(response.Status, 6, 12)
	fmt.Printf("Status %d ", status)
	if err != nil {
		fmt.Print("Error response status")
	}
	if status > 400 {
		fmt.Printf("Status Code : %s")
	}

	fmt.Print(response.Body)
}
