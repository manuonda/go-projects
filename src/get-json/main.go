package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type person struct {
	Name string `json:"name"`
}
type people struct {
	Number int      `json:"number"`
	Person []person `json:"people"`
}

func main() {
	url := "http://api.open-notify.org/astros.json"
	people, err := getAstros(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d people found in space  .\n", people.Number)

	for _, p := range people.Person {
		fmt.Printf("Let's wate to: %s\n", p.Name)
	}
}

func getAstros(apiUrl string) (people, error) {
	p := people{}
	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	// Create NewRequest
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return p, err
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, err := spaceClient.Do(req)

	if err != nil {
		return p, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Read All
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatalf("unable to parse value %q , error :%s ", string(body), err.Error())
	}

	return p, nil
}

func main2() {
	fmt.Println("Hello world")
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS",
	"name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft":
	"ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft":
	"ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`

	// convert to bytes
	textBytes := []byte(text)

	// create type of people
	p := people{}
	//json.Unmarshall functon works with a []byte type instead of a string uset []byte
	// to create the format we need
	err := json.Unmarshal(textBytes, &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p.Number)

}
