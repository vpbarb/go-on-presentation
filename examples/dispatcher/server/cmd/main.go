package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher"
	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/server"
)

const (
	workersCount = 3
	sendInterval = 10 * time.Second
	queueSize    = 10000
	maxBatchSize = 10
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	// START1 OMIT
	d := dispatcher.New(workersCount, sendInterval, queueSize, maxBatchSize)

	signalChan := make(chan os.Signal, 1)            // HL
	signal.Notify(signalChan, os.Interrupt, os.Kill) // HL

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(listener, server.NewHandler(d))

	log.Printf("Start listening on %s", listener.Addr())

	stopChan := make(chan struct{}) // HL

	go func() {
		<-signalChan // HL
		listener.Close()
		close(stopChan)
	}()

	d.Run(stopChan)
	// STOP1 OMIT
}
