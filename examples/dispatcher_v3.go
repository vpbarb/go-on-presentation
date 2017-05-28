package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v3"
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
		SendInterval: 300 * time.Millisecond, // HL
	}
	d := dispatcher.New(cfg)
	go d.Run()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Fake delay HL
		d.Add(dispatcher.Payload{"key": fmt.Sprintf("value_%d", i)})
	}
	time.Sleep(100 * time.Millisecond)
	// STOP1 OMIT
}
