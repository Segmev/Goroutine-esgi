package main

import (
	"fmt"
	"time"
)

func inRoutine() {
	fmt.Println("I'm printing from a routine")
}

func main() {
	go inRoutine()
	fmt.Println("I'm printing from the main process")
	time.Sleep(1 * time.Second)
	fmt.Println("I'm printing from the main process again after 1 second")
}
