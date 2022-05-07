package main

import "fmt"

func main() {
	channel := make(chan string, 2 /* this determines the size of the channel, or, how much data it can buffer within one instance*/)
	// buffer prevents deadlock

	channel <- "Buffer1"
	channel <- "Buffer2"

	msg := <-channel

	msg2 := <-channel

	fmt.Println(msg, msg2)
}
