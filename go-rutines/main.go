package main

import "fmt"

func main() {
	numberChannel := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numberChannel <- i
		}
		close(numberChannel)
	}()

	for number := range numberChannel {
		fmt.Println(number)
	}
}
