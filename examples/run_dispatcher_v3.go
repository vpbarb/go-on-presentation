package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v3"
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
		FlushInterval: 300 * time.Millisecond, // HL
	}
	d := dispatcher.New(cfg, &processor.Fake{})
	d.Run()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Fake delay // HL
		d.Collect(dispatcher.Payload{"value": fmt.Sprintf("%d", i)})
	}
	time.Sleep(500 * time.Millisecond) // HL
}

// STOP OMIT
