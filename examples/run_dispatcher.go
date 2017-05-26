package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	// START1 OMIT
	d := dispatcher.New(2, time.Second, 1000, 10)
	stopChan := make(chan struct{})
	go func() {
		for i := 0; i < 30; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			log.Print("add")
			d.Add(dispatcher.Payload{})
		}
		close(stopChan)
	}()
	d.Run(stopChan)
	// STOP1 OMIT
}
