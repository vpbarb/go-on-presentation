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
	cfg := collector.Config{
		MaxBatchSize:  3,
		WorkersCount:  3,
		QueueSize:     1000,
		FlushInterval: 300 * time.Millisecond, // HL
	}
	collector := collector.New(cfg, &processor.Fake{})
	collector.Run()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Fake delay // HL
		collector.Collect([]byte(fmt.Sprintf("event_%d", i)))
	}
	time.Sleep(300 * time.Millisecond) // HL
}

// STOP OMIT
