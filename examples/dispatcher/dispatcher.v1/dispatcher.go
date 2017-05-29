package dispatcher

import (
	"log"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

// START1 OMIT
type (
	Processor interface {
		Process(processor.Batch)
	}
	Payload    map[string]string
	Dispatcher struct {
		cfg       Config
		processor Processor
		batch     processor.Batch
	}
	Config struct {
		MaxBatchSize int
	}
)

func New(cfg Config, processor Processor) *Dispatcher {
	return &Dispatcher{cfg: cfg, processor: processor}
}

// STOP1 OMIT

// START2 OMIT
func (d *Dispatcher) Collect(payload Payload) error {
	d.batch = append(d.batch, processor.Item(payload))
	log.Printf("collected: %v", payload)

	if len(d.batch) >= d.cfg.MaxBatchSize { // HL
		t := time.Now()
		d.processor.Process(d.batch) // HL
		log.Printf("flushed %d payloads in %s", len(d.batch), time.Since(t))
		d.batch = nil // HL
	}

	return nil
}

// STOP2 OMIT
