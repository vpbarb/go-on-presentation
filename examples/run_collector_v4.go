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
	collector := &collector.Collector{
		Processor:     &processor.Fake{},
		MaxBatchSize:  3,
		WorkersCount:  3,
		QueueSize:     1000,
		FlushInterval: 300 * time.Millisecond,
	}

	stopChan := make(chan struct{})

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			collector.Collect(fmt.Sprintf("event_%d", i))
		}
		close(stopChan)
	}()

	collector.Run(stopChan)

	err := collector.Collect("slowpoke")
	log.Printf("collection error: %v", err)
}
