package main

import "fmt"

func main() {
	// START OMIT
	messages := make(chan string, 100)
	go func() { messages <- "iteration"; close(messages) }()
	for message := range messages {
		fmt.Println(message)
	}
	// STOP OMIT
}
