package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "channel1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "channel2"
		}
	}()

	for {
		select {
		case msgChannel1 := <-channel1:
			fmt.Println(msgChannel1)
		case msgChannel2 := <-channel2:
			fmt.Println(msgChannel2)
		}
		// with select execution of one channel does not need to wait for the other
		// msgChannel1 := <-channel1
		// fmt.Println(msgChannel1)

		// msgChannel2 := <-channel2
		// fmt.Println(msgChannel2)
	}
}
