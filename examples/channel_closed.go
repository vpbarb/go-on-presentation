package main

func main() {
	// START OMIT
	messages := make(chan string)
	close(messages)
	messages <- "ping" // Will panic
	// STOP OMIT
}
