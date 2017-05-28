package main

import (
	"fmt"
	"log"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v1"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// START1 OMIT
	cfg := dispatcher.Config{
		MaxBatchSize: 3, // HL
	}
	d := dispatcher.New(cfg)
	for i := 1; i <= 10; i++ {
		d.Add(dispatcher.Payload{"key": fmt.Sprintf("value_%d", i)}) // HL
	}
	// STOP1 OMIT
}
