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
	log.Print("dispatcher start")
	defer log.Print("dispatcher stop")

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

func (d *Dispatcher) Add(p Payload) {
	d.queue <- p
}

func (d *Dispatcher) worker(i int) {
	var batch []Payload

	log.Printf("[wrk_%d] start", i)
	defer log.Printf("[wrk_%d] stop", i)

	timer := time.NewTimer(d.sendInterval)
	defer timer.Stop()

	send := func(reason string) {
		if len(batch) == 0 {
			return
		}
		t := time.Now()
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Fake delay for example
		log.Printf("[wrk_%d, %s] sent %d payloads in %s", i, reason, len(batch), time.Since(t))
		batch = nil
		timer.Reset(d.sendInterval)
	}

	for {
		select {
		case p := <-d.queue:
			batch = append(batch, p)
			if len(batch) >= d.maxBatchSize {
				send("size")
			}
		case <-timer.C:
			send("timer")
		case <-d.workersStop:
			send("stop")
			return
		}
	}
}

func (d *Dispatcher) send(workerIndex int, batch []Payload, reason string) {
	if len(batch) == 0 {
		return
	}
	t := time.Now()
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Fake delay for example
	log.Printf("[wrk_%d, %s] sent %d payloads in %s", workerIndex, reason, len(batch), time.Since(t))
}
