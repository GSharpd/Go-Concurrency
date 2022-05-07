package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(4)

	go func() {
		write("Routine 1")
		waitgroup.Done() // removes 1 from quantity of go routines in waitgroup
	}()

	go func() {
		write("Routine 2")
		waitgroup.Done() // -1
	}()

	go func() {
		write("Routine 3")
		waitgroup.Done() // -1
	}()

	go func() {
		write("Routine 4")
		waitgroup.Done() // -1
	}()

	waitgroup.Wait() // waits until routines get to 0
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
