package main

import "fmt"

func main() {
	// START OMIT
	messages := make(chan string)
	go func() { messages <- "ping" }()
	fmt.Println(<-messages)
	// STOP OMIT
}
