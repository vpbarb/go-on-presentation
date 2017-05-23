package dispatcher

import (
	"log"
	"sync"
	"time"
)

type (
	Dispatcher struct {
		workersCount int
		sendInterval time.Duration
		maxBatchSize int
		done         chan struct{}
		queue        chan event
		workersWG    sync.WaitGroup
	}

	event map[string]string
)

func New(workersCount int, sendInterval time.Duration, bufSize int, maxBatchSize int) *Dispatcher {
	return &Dispatcher{
		workersCount: workersCount,
		sendInterval: sendInterval,
		maxBatchSize: maxBatchSize,
		done:         make(chan struct{}),
		queue:        make(chan event, bufSize),
	}
}

func (d *Dispatcher) Start() {
	d.workersWG.Add(d.workersCount)
	for i := 0; i < d.workersCount; i++ {
		go func(idx int) {
			defer d.workersWG.Done()
			d.worker(idx)
		}(i)
	}
	log.Println("dispatcher is started")
}

func (d *Dispatcher) Stop() {
	close(d.done)
	d.workersWG.Wait()
	log.Println("dispatcher is stopped")
}

func (d *Dispatcher) Add(event event) {
	d.queue <- event
}

func (d *Dispatcher) worker(idx int) {
	var batch []event

	log.Printf("worker %d is started\n", idx)
	defer log.Printf("worker %d is stopped\n", idx)

	send := func(batch []event) {
		if len(batch) == 0 {
			return
		}
		log.Printf("start sending batch of %d events\n", len(batch))
		time.Sleep(250 * time.Millisecond) // Fake delay to show example
		log.Printf("complete sending batch of %d events\n", len(batch))
	}

	for {
		select {
		case event := <-d.queue:
			batch = append(batch, event)
			if len(batch) >= d.maxBatchSize {
				send(batch)
				batch = nil
			}
		case <-time.After(d.sendInterval):
			send(batch)
			batch = nil
		case <-d.done:
			return
		}
	}
}
