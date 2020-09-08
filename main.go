package main

import (
	"fmt"
)

func inRoutine(channel chan int) {
	fmt.Println("I'm printing from a routine")
	channel <- 0
}

func main() {
	channel := make(chan int)

	go inRoutine(channel)
	sent := <-channel
	fmt.Println("I'm printing from the main process with", sent)
}
