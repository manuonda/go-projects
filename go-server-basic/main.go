package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHanlder(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "Address = %s", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "methos is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHanlder)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server starting at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening at port 8080")
}
