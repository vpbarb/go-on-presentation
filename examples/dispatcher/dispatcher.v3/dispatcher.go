package dispatcher

import (
	"log"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

type (
	Processor interface {
		Process(processor.Batch)
	}
	Payload map[string]string
)

// START1 OMIT
type (
	Dispatcher struct {
		cfg       Config
		processor Processor
		queue     chan Payload
	}
	Config struct {
		WorkersCount  int
		MaxBatchSize  int
		QueueSize     int
		FlushInterval time.Duration // HL
	}
)

// STOP1 OMIT

func New(cfg Config, processor Processor) *Dispatcher {
	return &Dispatcher{
		processor: processor,
		cfg:       cfg,
	}
}

func (d *Dispatcher) Add(payload Payload) error {
	log.Printf("add payload: %v", payload)
	d.queue <- payload
	return nil
}

func (d *Dispatcher) Run() {
	log.Print("dispatcher start")

	d.queue = make(chan Payload, d.cfg.QueueSize)

	for i := 0; i < d.cfg.WorkersCount; i++ {
		go func(i int) {
			d.worker(i)
		}(i)
	}
}

// START2 OMIT
func (d *Dispatcher) worker(i int) {
	var batch processor.Batch

	log.Printf("wrk_%d start", i)

	timer := time.NewTimer(d.cfg.FlushInterval) // HL

	flush := func(reason string) {
		t := time.Now()
		d.processor.Process(batch)
		log.Printf("wrk_%d flushed by '%s' %d payloads in %s", i, reason, len(batch), time.Since(t))
		batch = nil
		timer.Reset(d.cfg.FlushInterval) // HL
	}

	for {
		select { // HL
		case payload := <-d.queue: // HL
			batch = append(batch, processor.Item(payload))
			if len(batch) >= d.cfg.MaxBatchSize {
				flush("size")
			}
		case <-timer.C: // HL
			flush("timer")
		} // HL
	}
}

// STOP2 OMIT
