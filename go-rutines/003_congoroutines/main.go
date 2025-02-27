package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Goroutines: Son funciones o métodos que se ejecutan de manera concurrente.
sync.WaitGroup: Permite esperar a que un grupo de goroutines termine antes de continuar.
*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // indica que la goroutine a terminado
	fmt.Printf("Trabajador %d iniciando\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Trabajador %d finalizando\n", id)

}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(i)
		go worker(i, &wg)
	}

	wg.Wait() //Espera hasta que todas las goroutines terminen
	fmt.Println("Todos los trabajos han terminado")

}
