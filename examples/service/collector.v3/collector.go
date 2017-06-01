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
		WorkersCount int
		MaxBatchSize int
		QueueSize    int

		FlushInterval time.Duration // HL

		payloadsQueue chan string
	}
)

// STOP1 OMIT

func (c *Collector) Collect(payload string) {
	c.payloadsQueue <- payload
	log.Printf("collected: %s", payload)
}

func (c *Collector) Run() {
	log.Print("collector start")

	c.payloadsQueue = make(chan string, c.QueueSize)

	for i := 0; i < c.WorkersCount; i++ {
		go func(id int) {
			c.worker(id)
		}(i)
	}
}

// START2 OMIT
func (c *Collector) worker(id int) {
	var batch processor.Batch

	log.Printf("worker_%d start", id) // OMIT
	// OMIT
	timer := time.NewTimer(c.FlushInterval) // HL

	for {
		select { // HL
		case payload := <-c.payloadsQueue: // HL
			batch = append(batch, payload)
			if len(batch) >= c.MaxBatchSize {
				c.flush(id, batch, "size")
				batch = nil
				timer.Reset(c.FlushInterval) // HL
			}
		case <-timer.C: // HL
			c.flush(id, batch, "timer")
			batch = nil
			timer.Reset(c.FlushInterval) // HL
		} // HL
	}
}

// STOP2 OMIT
func (c *Collector) flush(workerId int, batch processor.Batch, reason string) {
	t := time.Now()
	c.Processor.Process(batch) // HL
	log.Printf("worker_%d flushed %d payloads by '%s' in %s", workerId, len(batch), reason, time.Since(t))
}
