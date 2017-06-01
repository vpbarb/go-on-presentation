package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/service/collector.v3"
	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
)

func init() {
	log.SetFlags(0)
}

// START OMIT
func main() {
	collector := &collector.Collector{
		Processor:     &processor.Fake{},
		MaxBatchSize:  2,
		WorkersCount:  2,
		QueueSize:     1000,
		FlushInterval: 200 * time.Millisecond, // HL
	}

	collector.Run()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Fake delay // HL
		collector.Collect(fmt.Sprintf("event_%d", i))
	}

	time.Sleep(300 * time.Millisecond)
}

// STOP OMIT
