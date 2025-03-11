package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	maxWorkers := 6
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	taskChan := make(chan int, len(tasks))
	resultChan := make(chan bool, len(tasks))
	wg := &sync.WaitGroup{}

	for i := 0; i < maxWorkers; i++ {
		go worker(taskChan, resultChan, wg)
	}

	//send task to workers
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	go func() {
		wg.Wait()
		//Cuando se cierra un canal, se pueden leer los datos
		//envaidos pero no enviar nuevos datos
		close(resultChan)
	}()

	for r := range resultChan {
		fmt.Println("resultado : ", r)
	}

}

// cuando se envia taskCHan chan: permite write/read in the same channel
// cuando es taskChan <- chan : only read
// cuando es taskChan chan<- :  only write
func worker(taskChan <-chan int, resultChan chan bool, wg *sync.WaitGroup) {

	for task := range taskChan {
		time.Sleep(time.Second)
		if task%3 == 0 {
			resultChan <- true
		} else {
			resultChan <- false
		}

		wg.Done()
	}
}
