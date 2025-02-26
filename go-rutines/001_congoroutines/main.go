package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	begin := time.Now()

	fmt.Println("==== Construction begins(Con rutines) ====")

	go Worker("Bob", "installing a door")
	go Worker("Alice", "mowing the lawn")
	go Worker("John", "painting the walls")

	//le agrego que espere el main
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf(" ==== Construction ended in %2.f seconds ====\n", time.Since(begin).Seconds())
}

func Worker(name, job string) {

	fmt.Printf("Worker %s began to work on %s\n", name, job)

	for i := 0; i < 10; i++ {
		fmt.Printf("Worker %s is working (%d/%d)..\n",
			name, i+1, 10)
		time.Sleep(50 * time.Millisecond)
	}
}
