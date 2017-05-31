package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/service/collector.v2"
	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
)

func init() {
	log.SetFlags(0)
}

// START OMIT
func main() {
	collector := &collector.Collector{
		Processor:    &processor.Fake{},
		MaxBatchSize: 2,
		WorkersCount: 2,    // HL
		QueueSize:    1000, // HL
	}

	collector.Run() // HL

	for i := 1; i <= 5; i++ {
		collector.Collect([]byte(fmt.Sprintf("event_%d", i)))
	}

	time.Sleep(200 * time.Millisecond)
}

// STOP OMIT
