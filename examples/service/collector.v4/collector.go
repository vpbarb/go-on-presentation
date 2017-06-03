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
		payloadsQueue chan string
	}
)

func (c *Collector) Collect(payload string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("collector is not running")
		}
	}()
	c.payloadsQueue <- payload
	log.Printf("collected: %s", payload)
	return nil
}

func (c *Collector) Run(stop chan struct{}) {
	log.Print("collector start")
	defer log.Print("collector stop")

	c.payloadsQueue = make(chan string, c.QueueSize)

	wg := sync.WaitGroup{}
	wg.Add(c.WorkersCount)
	for i := 0; i < c.WorkersCount; i++ {
		go func(i int) {
			defer wg.Done()
			c.worker(i)
		}(i)
	}

	<-stop
	close(c.payloadsQueue)
	wg.Wait()
}

func (c *Collector) worker(id int) {
	var buffer processor.Batch

	log.Printf("worker_%d start", id)
	defer log.Printf("worker_%d stop", id)

	timer := time.NewTimer(c.FlushInterval)
	defer timer.Stop()

	for {
		select {
		case payload, opened := <-c.payloadsQueue:
			if !opened {
				c.flush(id, buffer, "stop")
				return
			}
			buffer = append(buffer, payload)
			if len(buffer) >= c.MaxBatchSize {
				c.flush(id, buffer, "size")
				buffer = nil
				timer.Reset(c.FlushInterval)
			}
		case <-timer.C:
			c.flush(id, buffer, "timer")
			buffer = nil
			timer.Reset(c.FlushInterval)
		}
	}
}

func (c *Collector) flush(workerId int, batch processor.Batch, reason string) {
	t := time.Now()
	c.Processor.Process(batch)
	log.Printf("worker_%d flushed %d payloads by '%s' in %s", workerId, len(batch), reason, time.Since(t))
}
