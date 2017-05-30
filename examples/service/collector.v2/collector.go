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
		queue     chan []byte // HL
	}
	Config struct {
		MaxBatchSize int
		WorkersCount int // HL
		QueueSize    int // HL
	}
)

// STOP1 OMIT

func New(cfg Config, processor Processor) *Collector {
	return &Collector{
		cfg:       cfg,
		processor: processor,
	}
}

// START2 OMIT
func (c *Collector) Collect(payload []byte) error {
	c.queue <- payload // HL
	log.Printf("collected: %s", payload)
	return nil
}

func (c *Collector) Run() {
	log.Print("collector start")

	c.queue = make(chan []byte, c.cfg.QueueSize) // HL

	for i := 0; i < c.cfg.WorkersCount; i++ { // HL
		go func(i int) { // HL
			c.worker(i) // HL
		}(i) // HL
	} // HL
}

// STOP2 OMIT

// START3 OMIT
func (c *Collector) worker(i int) {
	var batch processor.Batch // HL

	log.Printf("worker_%d start", i)

	flush := func() { // HL
		t := time.Now()
		c.processor.Process(batch) // HL
		log.Printf("worker_%d flushed %d payloads in %s", i, len(batch), time.Since(t))
		batch = nil // HL
	} // HL

	for payload := range c.queue { // HL
		batch = append(batch, payload)
		if len(batch) >= c.cfg.MaxBatchSize {
			flush() // HL
		}
	} // HL
}

// STOP3 OMIT
