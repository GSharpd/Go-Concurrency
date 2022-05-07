package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("1")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received: %s", text)
			time.Sleep(time.Second / 2)
		}
	}()

	return channel
}
