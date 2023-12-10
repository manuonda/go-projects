package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello world")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, "Hello World")
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "The Godfather: Part II", Director: "Francis Ford Coppola"},
			},
		}
		tmpl.Execute(w, films)

	}
	http.HandleFunc("/", h1)

	handlePost := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		log.Print("HTMX request received")
		log.Print(r.Header.Get("hx-request"))
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		fmt.Println(title)
		fmt.Print(director)
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'><h4>%s</h4><p>Director: %s</p></li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/add-film/", handlePost)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
