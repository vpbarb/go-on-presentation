package main

import "fmt"

func main() {
	// START OMIT
	messages := make(chan string)

	messages <- "ping" // Block forever

	fmt.Println(<-messages)
	// STOP OMIT
}
