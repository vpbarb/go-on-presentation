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

		payloadsQueue chan string // HL
	}
)

// STOP1 OMIT

// START2 OMIT
func (c *Collector) Collect(payload string) {
	c.payloadsQueue <- payload // HL
	log.Printf("collected: %s", payload)
}

func (c *Collector) Run() { // HL
	log.Print("collector start")

	c.payloadsQueue = make(chan string, c.QueueSize) // HL

	for i := 0; i < c.WorkersCount; i++ {
		go func(id int) { // HL
			c.worker(id) // HL
		}(i) // HL
	}
} // HL

// STOP2 OMIT

// START3 OMIT
func (c *Collector) worker(id int) {
	var buffer processor.Batch // HL

	log.Printf("worker_%d start", id)

	for payload := range c.payloadsQueue { // HL
		buffer = append(buffer, payload)
		if len(buffer) >= c.MaxBatchSize {
			c.flush(id, buffer)
			buffer = nil
		}
	} // HL
}

// STOP3 OMIT
func (c *Collector) flush(workerId int, batch processor.Batch) {
	t := time.Now()
	c.Processor.Process(batch) // HL
	log.Printf("worker_%d flushed %d payloads in %s", workerId, len(batch), time.Since(t))
}
