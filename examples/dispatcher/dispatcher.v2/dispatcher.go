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
		queue     chan Payload // HL
	}
	Config struct {
		MaxBatchSize int
		WorkersCount int // HL
		QueueSize    int // HL
	}
)

// STOP1 OMIT

func New(cfg Config, processor Processor) *Dispatcher {
	return &Dispatcher{
		cfg:       cfg,
		processor: processor,
	}
}

// START2 OMIT
func (d *Dispatcher) Add(payload Payload) error {
	log.Printf("add payload: %v", payload)
	d.queue <- payload // HL
	return nil
}

func (d *Dispatcher) Run() {
	log.Print("dispatcher start")

	d.queue = make(chan Payload, d.cfg.QueueSize) // HL

	for i := 0; i < d.cfg.WorkersCount; i++ { // HL
		go func(i int) { // HL
			d.worker(i) // HL
		}(i) // HL
	} // HL
}

// STOP2 OMIT

// START3 OMIT
func (d *Dispatcher) worker(i int) {
	var batch processor.Batch // HL

	log.Printf("wrk_%d start", i)

	flush := func() { // HL
		t := time.Now()
		d.processor.Process(batch) // HL
		log.Printf("wrk_%d flushed %d payloads in %s", i, len(batch), time.Since(t))
		batch = nil // HL
	} // HL

	for payload := range d.queue { // HL
		batch = append(batch, processor.Item(payload))
		if len(batch) >= d.cfg.MaxBatchSize {
			flush() // HL
		}
	} // HL
}

// STOP3 OMIT
