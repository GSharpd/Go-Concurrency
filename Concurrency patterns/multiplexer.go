package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write("1"), write("2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplex(entryChannel1, entryChannel2 <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		for {
			select {
			case msg := <-entryChannel1:
				outChannel <- msg
			case msg := <-entryChannel2:
				outChannel <- msg
			}
		}
	}()

	return outChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
