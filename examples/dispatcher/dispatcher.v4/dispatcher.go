package dispatcher

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

type (
	Processor interface {
		Process(processor.Batch) error
	}
	Payload map[string]string

	Dispatcher struct {
		processor    Processor
		workersCount int
		maxBatchSize int
		queueSize    int
		sendInterval time.Duration
		queue        chan Payload
		isRun        bool           // HL
		workersWG    sync.WaitGroup // HL
	}
	Config struct {
		MaxBatchSize int
		WorkersCount int
		QueueSize    int
		SendInterval time.Duration
	}
)

func New(cfg Config) *Dispatcher {
	return &Dispatcher{
		processor:    &processor.Fake{},
		maxBatchSize: cfg.MaxBatchSize,
		workersCount: cfg.WorkersCount,
		queueSize:    cfg.QueueSize,
		sendInterval: cfg.SendInterval,
	}
}

func (d *Dispatcher) Add(payload Payload) error {
	if !d.isRun { // HL
		return errors.New("dispatcher is not running") // HL
	} // HL
	log.Printf("add payload: %v", payload)
	d.queue <- payload
	return nil
}

func (d *Dispatcher) Run(stopChan chan struct{}) {
	log.Print("dispatcher start")
	defer log.Print("dispatcher stop") // HL

	d.queue, d.isRun = make(chan Payload, d.queueSize), true // HL

	d.workersWG.Add(d.workersCount) // HL
	for i := 0; i < d.workersCount; i++ {
		go func(i int) {
			defer d.workersWG.Done() // HL
			d.worker(i)
		}(i)
	}

	<-stopChan
	close(d.queue)     // HL
	d.workersWG.Wait() // HL
}

func (d *Dispatcher) worker(i int) {
	var batch processor.Batch

	log.Printf("wrk_%d start", i)
	defer log.Printf("wrk_%d stop", i) // HL

	timer := time.NewTimer(d.sendInterval)
	defer timer.Stop() // HL

	flush := func(reason string) {
		t := time.Now()
		d.processor.Process(batch)
		log.Printf("wrk_%d proceed by '%s' %d payloads in %s", i, reason, len(batch), time.Since(t))
		batch = nil
		timer.Reset(d.sendInterval)
	}

	for {
		select {
		case payload, opened := <-d.queue: // HL
			if !opened { // HL
				flush("stop") // HL
				return        // HL
			} // HL
			batch = append(batch, processor.Item(payload))
			if len(batch) >= d.maxBatchSize {
				flush("size")
			}
		case <-timer.C:
			flush("timer")
		}
	}
}
