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
	// START1 OMIT
	Dispatcher struct {
		processor    Processor
		maxBatchSize int
		workersCount int          // HL
		queueSize    int          // HL
		queue        chan Payload // HL
	}
	Config struct {
		MaxBatchSize int
		WorkersCount int // HL
		QueueSize    int // HL
	}
	// STOP1 OMIT
)

func New(cfg Config) *Dispatcher {
	return &Dispatcher{
		processor:    &processor.Fake{},
		maxBatchSize: cfg.MaxBatchSize,
		workersCount: cfg.WorkersCount,
		queueSize:    cfg.QueueSize,
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

	d.queue = make(chan Payload, d.queueSize) // HL

	for i := 0; i < d.workersCount; i++ { // HL
		go func(i int) { // HL
			d.worker(i) // HL
		}(i) // HL
	} // HL
}

// STOP2 OMIT

// START3 OMIT
func (d *Dispatcher) worker(i int) {
	var batch processor.Batch

	log.Printf("wrk_%d start", i)

	flush := func() {
		t := time.Now()
		d.processor.Process(batch)
		log.Printf("wrk_%d proceed %d payloads in %s", i, len(batch), time.Since(t))
		batch = nil
	}

	for payload := range d.queue {
		batch = append(batch, processor.Item(payload))
		if len(batch) >= d.maxBatchSize {
			flush()
		}
	}
}

// STOP3 OMIT
