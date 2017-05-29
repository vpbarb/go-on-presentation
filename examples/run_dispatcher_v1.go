package main

import (
	"fmt"
	"log"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v1"
	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

func init() {
	log.SetFlags(0)
}

// START OMIT
func main() {
	cfg := dispatcher.Config{
		MaxBatchSize: 2, // HL
	}
	d := dispatcher.New(cfg, &processor.Fake{})
	for i := 1; i <= 5; i++ {
		d.Collect(dispatcher.Payload{"value": fmt.Sprintf("%d", i)}) // HL
	}
}

// STOP OMIT
