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
		Process(processor.Batch)
	}
	Payload map[string]string
)

type (
	Dispatcher struct {
		cfg       Config
		processor Processor
		queue     chan Payload
	}
	Config struct {
		MaxBatchSize  int
		WorkersCount  int
		QueueSize     int
		FlushInterval time.Duration
	}
)

func New(cfg Config, processor Processor) *Dispatcher {
	return &Dispatcher{
		cfg:       cfg,
		processor: processor,
	}
}

// START2 OMIT
func (d *Dispatcher) Add(payload Payload) (err error) {
	defer func() { // HL
		if r := recover(); r != nil { // HL
			err = errors.New("dispatcher is not running") // HL
		} // HL
	}() // HL
	d.queue <- payload
	log.Printf("added: %v", payload)
	return nil
}

// STOP2 OMIT

// START3 OMIT
func (d *Dispatcher) Run(stopChan chan struct{}) {
	log.Print("dispatcher start")
	defer log.Print("dispatcher stop") // HL

	d.queue = make(chan Payload, d.cfg.QueueSize) // HL
	wg := sync.WaitGroup{}                        // HL
	wg.Add(d.cfg.WorkersCount)                    // HL
	for i := 0; i < d.cfg.WorkersCount; i++ {
		go func(i int) {
			defer wg.Done() // HL
			d.worker(i)
		}(i)
	}

	<-stopChan     // HL
	close(d.queue) // HL
	wg.Wait()      // HL
}

// STOP3 OMIT

func (d *Dispatcher) worker(i int) {
	var batch processor.Batch

	log.Printf("wrk_%d start", i)
	defer log.Printf("wrk_%d stop", i) // HL

	timer := time.NewTimer(d.cfg.FlushInterval)
	defer timer.Stop() // HL

	flush := func(reason string) {
		t := time.Now()
		d.processor.Process(batch)
		log.Printf("wrk_%d flushed by '%s' %d payloads in %s", i, reason, len(batch), time.Since(t))
		batch = nil
		timer.Reset(d.cfg.FlushInterval)
	}

	// START4 OMIT
	for {
		select {
		case payload, opened := <-d.queue: // HL
			if !opened { // HL
				flush("stop") // HL
				return        // HL
			} // HL
			batch = append(batch, processor.Item(payload))
			if len(batch) >= d.cfg.MaxBatchSize {
				flush("size")
			}
		case <-timer.C:
			flush("timer")
		}
	}
	// STOP4 OMIT
}
