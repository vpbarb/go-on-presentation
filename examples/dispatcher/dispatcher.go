package dispatcher

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type (
	Dispatcher struct {
		workersCount int
		sendInterval time.Duration
		maxBatchSize int
		queue        chan Payload
		workersStop  chan struct{}
		workersWG    sync.WaitGroup
	}

	Payload map[string]string
)

func New(workersCount int, sendInterval time.Duration, queueSize int, maxBatchSize int) *Dispatcher {
	return &Dispatcher{
		workersCount: workersCount,
		sendInterval: sendInterval,
		maxBatchSize: maxBatchSize,
		queue:        make(chan Payload, queueSize),
		workersStop:  make(chan struct{}),
	}
}

func (d *Dispatcher) Run(stop chan struct{}) {
	log.Print("dispatcher is started")
	defer log.Print("dispatcher is stopped")

	d.workersWG.Add(d.workersCount)
	for i := 0; i < d.workersCount; i++ {
		go func(i int) {
			defer d.workersWG.Done()
			d.worker(i)
		}(i)
	}

	<-stop
	close(d.workersStop)
	d.workersWG.Wait()
}

func (d *Dispatcher) Send(p Payload) {
	d.queue <- p
}

func (d *Dispatcher) worker(i int) {
	var batch []Payload

	log.Printf("worker %d is started", i)
	defer log.Printf("worker %d is stopped", i)

	for {
		select {
		case p := <-d.queue:
			batch = append(batch, p)
			if len(batch) >= d.maxBatchSize {
				log.Printf("worker %d sends batch because of size", i)
				d.send(i, batch)
				batch = nil
			}
		case <-time.After(d.sendInterval):
			log.Printf("worker %d sends batch because of time", i)
			d.send(i, batch)
			batch = nil
		case <-d.workersStop:
			log.Printf("worker %d sends batch because of stop", i)
			d.send(i, batch)
			return
		}
	}
}

func (d *Dispatcher) send(workerIndex int, batch []Payload) {
	if len(batch) == 0 {
		return
	}
	t := time.Now()
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Fake delay for example
	log.Printf("worker %d sent batch of %d payloads in %s", workerIndex, len(batch), time.Since(t))
}
