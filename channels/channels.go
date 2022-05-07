package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Routine", channel)

	fmt.Println("After write() routine call")

	// channels can send a receive data
	for {
		msg, isOpen := <-channel // sends data to the variable
		if !isOpen {
			fmt.Println("Channel closed")
			break // checks if channel is open and breaks execution
		}
		fmt.Println("Channel is open")
		fmt.Println(msg)
	}

	// Range can be used with channels
	// for msg := range channel { fmt.Println(msg) } would also work

	fmt.Println("End of execution")
}

func write(text string, channel chan string) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		channel <- text // channel receives data from the parameter
		time.Sleep(time.Second)
	}
	fmt.Println("Closing channel")
	close(channel)
}
