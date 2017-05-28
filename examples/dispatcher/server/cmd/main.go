package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v4"
	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/server"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	// START1 OMIT
	cfg := dispatcher.Config{
		WorkersCount:  3,
		FlushInterval: 5 * time.Second,
		QueueSize:     10000,
		MaxBatchSize:  10,
	}
	d := dispatcher.New(cfg, &processor.Fake{})

	signalChan := make(chan os.Signal, 1)            // HL
	signal.Notify(signalChan, os.Interrupt, os.Kill) // HL

	listener, err := net.Listen("tcp", "localhost:9090") // HL
	if err != nil {                                      // HL
		log.Fatal(err) // HL
	} // HL

	go http.Serve(listener, server.NewDispatcherHandler(d)) // HL
	log.Printf("Start listening on %s", listener.Addr())    // OMIT

	stopChan := make(chan struct{})
	go func() {
		<-signalChan     // HL
		listener.Close() // HL
		close(stopChan)
	}()
	d.Run(stopChan)
	// STOP1 OMIT
}
