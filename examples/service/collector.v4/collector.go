package collector

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
)

type (
	Processor interface {
		Process(processor.Batch)
	}
)

type (
	Collector struct {
		cfg       Config
		processor Processor
		queue     chan []byte
	}
	Config struct {
		MaxBatchSize  int
		WorkersCount  int
		QueueSize     int
		FlushInterval time.Duration
	}
)

func New(cfg Config, processor Processor) *Collector {
	return &Collector{
		cfg:       cfg,
		processor: processor,
	}
}

// START2 OMIT
func (c *Collector) Collect(payload []byte) (err error) {
	defer func() { // HL
		if r := recover(); r != nil { // HL
			err = errors.New("collector is not running") // HL
		} // HL
	}() // HL
	c.queue <- payload
	log.Printf("collected: %s", payload)
	return nil
}

// STOP2 OMIT

// START3 OMIT
func (c *Collector) Run(stop chan struct{}) {
	log.Print("collector start")
	defer log.Print("collector stop") // HL

	c.queue = make(chan []byte, c.cfg.QueueSize) // HL

	wg := sync.WaitGroup{}     // HL
	wg.Add(c.cfg.WorkersCount) // HL
	for i := 0; i < c.cfg.WorkersCount; i++ {
		go func(i int) {
			defer wg.Done() // HL
			c.worker(i)
		}(i)
	}

	<-stop         // HL
	close(c.queue) // HL
	wg.Wait()      // HL
}

// STOP3 OMIT

func (c *Collector) worker(i int) {
	var batch processor.Batch

	log.Printf("worker_%d start", i)
	defer log.Printf("worker_%d stop", i)

	timer := time.NewTimer(c.cfg.FlushInterval)
	defer timer.Stop()

	flush := func(reason string) {
		t := time.Now()
		c.processor.Process(batch)
		log.Printf("worker_%d flushed by '%s' %d payloads in %s", i, reason, len(batch), time.Since(t))
		batch = nil
		timer.Reset(c.cfg.FlushInterval)
	}

	// START4 OMIT
	for {
		select {
		case payload, opened := <-c.queue: // HL
			if !opened { // HL
				flush("stop") // HL
				return        // HL
			} // HL
			batch = append(batch, payload)
			if len(batch) >= c.cfg.MaxBatchSize {
				flush("size")
			}
		case <-timer.C:
			flush("timer")
		}
	}
	// STOP4 OMIT
}
