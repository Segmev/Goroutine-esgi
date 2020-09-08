package main

import (
	"fmt"
)

func inRoutine(channel chan int, id int) {
	fmt.Println("I'm printing from a routine")
	channel <- id
}

func main() {
	channel := make(chan int)

	go inRoutine(channel, 0)
	go inRoutine(channel, 1)
	go inRoutine(channel, 2)
	sent, sent2, sent3 := <-channel, <-channel, <-channel
	fmt.Println("I'm printing from the main process with", sent, sent2, sent3)
}
