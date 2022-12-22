package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Counter %v", ct.counter)
	ct.counter++
	fmt.Fprintf(w, "Counter : %v ", ct.counter)
}

func main() {
	th := &CounterHandler{counter: 0}
	http.Handle("/count", th)
	http.ListenAndServe(":8080", nil)

}
