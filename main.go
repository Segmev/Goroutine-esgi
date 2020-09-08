package main

import (
	"fmt"
	"time"
)

func waitingRoutine(channelIn chan string, channelOut chan string) {
	fmt.Println("I'm the routine, I'll wait for a message")
	msg := <-channelIn
	fmt.Printf("Received this message:\n%s\n", msg)
	fmt.Printf(".\n.\n.\n")
	channelOut <- "Always been"
}

func main() {
	channelIn, channelOut := make(chan string), make(chan string)

	go waitingRoutine(channelIn, channelOut)

	fmt.Println("I'm the main process, I launched the routine and I'll wait some time")
	time.Sleep(3 * time.Second)
	fmt.Printf("I stopped waiting, sending a message...")
	channelIn <- "Wait, it's that simple ?"
	fmt.Println(<-channelOut)
}
