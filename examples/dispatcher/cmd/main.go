package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"
	"net"
	"log"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher"
)

const (
	workersCount     = 3
	sendInterval     = 5 * time.Second
	eventsBufferSize = 10000
	maxBatchSize     = 5
)

func main() {
	d := dispatcher.New(workersCount, sendInterval, eventsBufferSize, maxBatchSize)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(listener, dispatcher.NewEventHandler(d))

	d.Start()
	<-signalChan
	listener.Close()
	d.Stop()
}
