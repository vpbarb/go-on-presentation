package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v4"
	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

func init() {
	log.SetFlags(0)
}

// START OMIT
func main() {
	cfg := dispatcher.Config{
		MaxBatchSize:  3,
		WorkersCount:  3,
		QueueSize:     1000,
		FlushInterval: 300 * time.Millisecond,
	}
	d := dispatcher.New(cfg, &processor.Fake{})
	stopChan := make(chan struct{}) // HL
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Fake delay
			d.Add(dispatcher.Payload{"value": fmt.Sprintf("%d", i)})
		}
		close(stopChan) // HL
	}()
	go func() { // HL
		<-stopChan                                            // HL
		err := d.Add(dispatcher.Payload{"value": "slowpoke"}) // HL
		log.Printf("add error: %v", err)                      // HL
	}() // HL
	d.Run(stopChan) // HL
}

// STOP OMIT
