package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.RWMutex: Similar a Mutex,
              pero permite múltiples lectores simultáneamente
			  mientras no haya escritores.
**/

var data = 0
var rw sync.RWMutex

func read(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.RLock() //Bloque solo para lectura
	fmt.Printf("Lector %d : leyendo :%d\n", id, data)
	time.Sleep(time.Second)
	rw.RUnlock() //desbloquea solo par alectura
}

func write(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.Lock() //bloquea para escritura (nadie mas puede leer o escribir)
	data++
	fmt.Printf("Escritor :%d modifico el valor :%d\n", id, data)
	time.Sleep(time.Second)
	rw.Unlock()
}

func main() {
	var wg sync.WaitGroup

	//lanzamos 3 lectores
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go read(i, &wg)
	}

	//lanzamos un escritor
	wg.Add(1)
	go write(1, &wg)

	wg.Wait()

}
