package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v2"
	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

func init() {
	log.SetFlags(0)
}

// START OMIT
func main() {
	cfg := dispatcher.Config{
		MaxBatchSize: 2,
		WorkersCount: 2,    // HL
		QueueSize:    1000, // HL
	}
	d := dispatcher.New(cfg, &processor.Fake{})
	d.Run() // HL
	for i := 1; i <= 5; i++ {
		d.Collect(dispatcher.Payload{"value": fmt.Sprintf("%d", i)})
	}
	time.Sleep(100 * time.Millisecond) // HL
}

// STOP OMIT
