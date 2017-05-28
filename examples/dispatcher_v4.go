package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v4"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// START1 OMIT
	cfg := dispatcher.Config{
		MaxBatchSize: 3,
		WorkersCount: 3,
		QueueSize:    1000,
		SendInterval: 300 * time.Millisecond,
	}
	d := dispatcher.New(cfg)
	stopChan := make(chan struct{}) // HL
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Fake delay
			d.Add(dispatcher.Payload{"key": fmt.Sprintf("value_%d", i)})
		}
		close(stopChan) // HL
	}()
	d.Run(stopChan) // HL
	// STOP1 OMIT
}
