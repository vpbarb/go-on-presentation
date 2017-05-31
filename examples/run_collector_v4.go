package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/service/collector.v4"
	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// START OMIT
	collector := &collector.Collector{
		Processor:     &processor.Fake{},
		MaxBatchSize:  3,
		WorkersCount:  3,
		QueueSize:     1000,
		FlushInterval: 300 * time.Millisecond,
	}

	stopChan := make(chan struct{}) // HL

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Fake delay
			collector.Collect([]byte(fmt.Sprintf("event_%d", i)))
		}
		close(stopChan) // HL
	}()

	collector.Run(stopChan) // HL

	err := collector.Collect([]byte("slowpoke")) // HL
	log.Printf("collection error: %v", err)      // HL
	// STOP OMIT
}
