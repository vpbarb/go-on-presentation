package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/service/collector.v4"
	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
	"github.com/Barberrrry/go-on-presentation/examples/service/server"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

// START1 OMIT
func main() {
	cfg := collector.Config{WorkersCount: 3, FlushInterval: 5 * time.Second, QueueSize: 10000, MaxBatchSize: 10}
	collector := collector.New(cfg, &processor.Fake{})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(listener, server.NewCollectorHandler(collector))

	stopChan := make(chan struct{})
	go func() {
		<-signalChan
		listener.Close()
		close(stopChan)
	}()

	collector.Run(stopChan)
}

// STOP1 OMIT
