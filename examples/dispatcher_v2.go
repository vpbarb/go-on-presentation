package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v2"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// START1 OMIT
	cfg := dispatcher.Config{
		MaxBatchSize: 3,
		WorkersCount: 3,    // HL
		QueueSize:    1000, // HL
	}
	d := dispatcher.New(cfg)
	go d.Run() // HL
	for i := 1; i <= 10; i++ {
		d.Add(dispatcher.Payload{"key": fmt.Sprintf("value_%d", i)})
	}
	time.Sleep(100 * time.Millisecond) // HL
	// STOP1 OMIT
}
