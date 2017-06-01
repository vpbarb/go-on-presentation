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
	collector := &collector.Collector{
		Processor:    &processor.Fake{},
		MaxBatchSize: 2,
	}

	for i := 1; i <= 5; i++ {
		collector.Collect(fmt.Sprintf("event_%d", i))
	}
}

// STOP OMIT
