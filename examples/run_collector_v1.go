package main

import (
	"fmt"
	"log"

	"github.com/Barberrrry/go-on-presentation/examples/service/collector.v1"
	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
)

func init() {
	log.SetFlags(0)
}

// START OMIT
func main() {
	cfg := collector.Config{
		MaxBatchSize: 2, // HL
	}
	collector := collector.New(cfg, &processor.Fake{})
	for i := 1; i <= 5; i++ {
		collector.Collect([]byte(fmt.Sprintf("event_%d", i))) // HL
	}
}

// STOP OMIT
