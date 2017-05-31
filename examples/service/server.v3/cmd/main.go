package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/service/collector.v3"
	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
	"github.com/Barberrrry/go-on-presentation/examples/service/server.v3"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

// START1 OMIT
func main() {
	collector := &collector.Collector{
		Processor:     &processor.Fake{},
		MaxBatchSize:  5,
		WorkersCount:  3,
		QueueSize:     10000,
		FlushInterval: 5 * time.Second,
	}

	collector.Run() // HL

	http.ListenAndServe("localhost:9090", server.NewCollectorHandler(collector)) // HL
}

// STOP1 OMIT
