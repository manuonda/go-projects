package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var url = "https://saphirus.com.ar/vende-saphirus/"
	response, error := http.Get(url)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Print(response)
	status, err := strconv.ParseInt(response.Status, 6, 12)

	if err != nil {
		fmt.Print("E")
	}
	if status > 400 {
		fmt.Printf("Status Code : %s")
	}

	fmt.Print(response.Body)
}
