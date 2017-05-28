package dispatcher

import (
	"log"
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
		sendInterval time.Duration // HL
		queue        chan Payload
	}
	Config struct {
		WorkersCount int
		MaxBatchSize int
		QueueSize    int
		SendInterval time.Duration // HL
	}
)

func New(cfg Config) *Dispatcher {
	return &Dispatcher{
		processor:    &processor.Fake{},
		workersCount: cfg.WorkersCount,
		maxBatchSize: cfg.MaxBatchSize,
		queueSize:    cfg.QueueSize,
		sendInterval: cfg.SendInterval,
	}
}

func (d *Dispatcher) Add(payload Payload) error {
	log.Printf("add payload: %v", payload)
	d.queue <- payload
	return nil
}

func (d *Dispatcher) Run() {
	log.Print("dispatcher start")

	d.queue = make(chan Payload, d.queueSize)

	for i := 0; i < d.workersCount; i++ {
		go func(i int) {
			d.worker(i)
		}(i)
	}
}

func (d *Dispatcher) worker(i int) {
	var batch processor.Batch

	log.Printf("wrk_%d start", i)

	timer := time.NewTimer(d.sendInterval)

	flush := func(reason string) {
		t := time.Now()
		d.processor.Process(batch)
		log.Printf("wrk_%d proceed by '%s' %d payloads in %s", i, reason, len(batch), time.Since(t))
		batch = nil
		timer.Reset(d.sendInterval) // HL
	}

	for {
		select { // HL
		case payload := <-d.queue: // HL
			batch = append(batch, processor.Item(payload))
			if len(batch) >= d.maxBatchSize {
				flush("size")
			}
		case <-timer.C: // HL
			flush("timer")
		} // HL
	}
}
