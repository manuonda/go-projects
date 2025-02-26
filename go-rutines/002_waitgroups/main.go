package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	begin := time.Now()

	fmt.Println("==== Construction begins(Con rutines) ====")

	wg := &sync.WaitGroup{}

	//wg.Add(3)
	go Worker("Bob", "installing a door", wg)
	wg.Add(1)
	go Worker("Alice", "mowing the lawn", wg)
	wg.Add(1)
	go Worker("John", "painting the walls", wg)
	wg.Add(1)

	wg.Wait() //espero hasta que terminen los workers
	//3 indicaciones a terminar
	//le agrego que espere el main
	//time.Sleep(1000 * time.Millisecond)
	fmt.Printf(" ==== Construction ended in %2.f seconds ====\n", time.Since(begin).Seconds())
}

func Worker(name, job string, wg *sync.WaitGroup) {

	fmt.Printf("Worker %s began to work on %s\n", name, job)

	for i := 0; i < 10; i++ {
		fmt.Printf("Worker %s is working (%d/%d)..\n",
			name, i+1, 10)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Printf("Worker #{name} finished working on #{job}\n")
	wg.Done() // indico que finaliza el worker
}
