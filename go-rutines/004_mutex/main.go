package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Mutex: Se usa  para evitar condiciones de carrera
           cuando varias goroutines acceden a la misma
		   variable.
*/

var counter = 0

var mutex sync.Mutex

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("\n Job %d ", id)
	mutex.Lock()
	fmt.Printf("Trabajador %d iniciando \n", id)
	time.Sleep(time.Second)
	counter++
	fmt.Printf("Trabajador %d finalizando \n", id)
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	expectedCounter := 10

	fmt.Println("=== mutex =====")
	for i := 0; i < expectedCounter; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Todos los trabajos han terminado")
	// Check for race condition
	if expectedCounter != counter {
		fmt.Println("Race condition detected!")
	} else {
		fmt.Println("No race condition detected.")
	}
}
