package main

import (
	"fmt"
	"time"
)

func main() {
	// Concurrency != parallelism
	go write("Hello, world!")

	write("Go programming")

	// Never gets to second function call
	// call go before first function call to make it a go routine
	// meaning code execution wont wait for the call with go before it
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
