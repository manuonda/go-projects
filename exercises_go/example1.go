package main

import (
	"flag"
	"fmt"
)

/**
Example 1 - Build a CLI using the Go standard library “hello world!

This is a simple command-line interface (CLI) written in Go that outputs
 "Hello World" when a user enters a specific command.
 The program uses the flag package
from the Go standard library to define a command-line flag named message, which allows users to customize the output message.
**/

func main() {
	message := flag.String("message", "Star wars", "Print tmessage to stdout")
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println("Invalid arguments : ", flag.Args())
		return
	}
	fmt.Println(*message)
}
