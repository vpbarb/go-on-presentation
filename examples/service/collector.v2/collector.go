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
		Processor    Processor
		MaxBatchSize int
		WorkersCount int // HL
		QueueSize    int // HL

		queue chan []byte // HL
	}
)

// STOP1 OMIT

// START2 OMIT
func (c *Collector) Collect(payload []byte) {
	c.queue <- payload // HL
	log.Printf("collected: %s", payload)
}

func (c *Collector) Run() { // HL
	log.Print("collector start")

	c.queue = make(chan []byte, c.QueueSize) // HL

	for i := 0; i < c.WorkersCount; i++ {
		go func(id int) { // HL
			c.worker(id) // HL
		}(i) // HL
	}
} // HL

// STOP2 OMIT

// START3 OMIT
func (c *Collector) worker(id int) {
	var batch processor.Batch // HL

	log.Printf("worker_%d start", id)

	for payload := range c.queue { // HL
		batch = append(batch, payload)
		if len(batch) >= c.MaxBatchSize {
			c.flush(id, batch)
			batch = nil
		}
	} // HL
}

// STOP3 OMIT
func (c *Collector) flush(workerId int, batch processor.Batch) {
	t := time.Now()
	c.Processor.Process(batch) // HL
	log.Printf("worker_%d flushed %d payloads in %s", workerId, len(batch), time.Since(t))
}
