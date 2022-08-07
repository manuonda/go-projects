package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manuonda/go-projects/src/go-bookstore/pkg/routes"
)

func main() {
	fmt.Print("Init project")
	r := mux.NewRouter()
	fmt.Print("salida 1")
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("salida 2")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print("Error listentnt 8080")
	}
	fmt.Println("Listening port 8080")
}
