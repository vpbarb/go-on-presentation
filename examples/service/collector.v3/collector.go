package collector

import (
	"log"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/service/processor"
)

type (
	Processor interface {
		Process(processor.Batch)
	}
)

// START1 OMIT
type (
	Collector struct {
		cfg       Config
		processor Processor
		queue     chan []byte
	}
	Config struct {
		WorkersCount  int
		MaxBatchSize  int
		QueueSize     int
		FlushInterval time.Duration // HL
	}
)

// STOP1 OMIT

func New(cfg Config, processor Processor) *Collector {
	return &Collector{
		processor: processor,
		cfg:       cfg,
	}
}

func (c *Collector) Collect(payload []byte) error {
	c.queue <- payload
	log.Printf("collected: %s", payload)
	return nil
}

func (c *Collector) Run() {
	log.Print("collector start")

	c.queue = make(chan []byte, c.cfg.QueueSize)

	for i := 0; i < c.cfg.WorkersCount; i++ {
		go func(i int) {
			c.worker(i)
		}(i)
	}
}

// START2 OMIT
func (c *Collector) worker(i int) {
	var batch processor.Batch

	log.Printf("worker_%d start", i) // OMIT
	// OMIT
	timer := time.NewTimer(c.cfg.FlushInterval) // HL

	flush := func(reason string) {
		t := time.Now() // OMIT
		c.processor.Process(batch)
		log.Printf("worker_%d flushed by '%s' %d payloads in %s", i, reason, len(batch), time.Since(t)) // OMIT
		batch = nil
		timer.Reset(c.cfg.FlushInterval) // HL
	}

	for {
		select { // HL
		case payload := <-c.queue: // HL
			batch = append(batch, payload)
			if len(batch) >= c.cfg.MaxBatchSize {
				flush("size")
			}
		case <-timer.C: // HL
			flush("timer")
		} // HL
	}
}

// STOP2 OMIT
