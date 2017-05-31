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
		Processor     Processor
		MaxBatchSize  int
		WorkersCount  int
		QueueSize     int
		FlushInterval time.Duration
		queue         chan []byte
	}
)

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

	c.queue = make(chan []byte, c.QueueSize) // HL

	wg := sync.WaitGroup{} // HL
	wg.Add(c.WorkersCount) // HL
	for i := 0; i < c.WorkersCount; i++ {
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

func (c *Collector) worker(id int) {
	var batch processor.Batch

	log.Printf("worker_%d start", id)
	defer log.Printf("worker_%d stop", id)

	timer := time.NewTimer(c.FlushInterval)
	defer timer.Stop()

	// START4 OMIT
	for {
		select {
		case payload, opened := <-c.queue: // HL
			if !opened { // HL
				c.flush(id, batch, "stop") // HL
				return                     // HL
			} // HL
			batch = append(batch, payload)
			if len(batch) >= c.MaxBatchSize {
				c.flush(id, batch, "size")
				batch = nil
				timer.Reset(c.FlushInterval)
			}
		case <-timer.C:
			c.flush(id, batch, "timer")
			batch = nil
			timer.Reset(c.FlushInterval)
		}
	}
	// STOP4 OMIT
}

func (c *Collector) flush(workerId int, batch processor.Batch, reason string) {
	t := time.Now()
	c.Processor.Process(batch) // HL
	log.Printf("worker_%d flushed %d payloads in %s", workerId, len(batch), time.Since(t))
}
