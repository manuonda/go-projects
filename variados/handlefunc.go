package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloHandler)

	log.Println(("Listening"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

}
