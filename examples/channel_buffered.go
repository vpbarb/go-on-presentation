package main

import "fmt"

func main() {
	// START OMIT
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages, <-messages)
	// STOP OMIT
}
