package dispatcher

import (
	"log"
	"sync"
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
		flushLock sync.Mutex
		batch     processor.Batch
	}
	Config struct {
		MaxBatchSize int
	}
)

func New(cfg Config, processor Processor) *Dispatcher {
	return &Dispatcher{
		cfg:       cfg,
		processor: processor,
	}
}

// STOP1 OMIT

// START2 OMIT
func (d *Dispatcher) Add(payload Payload) error {
	d.flushLock.Lock()         // HL
	defer d.flushLock.Unlock() // HL

	d.batch = append(d.batch, processor.Item(payload))
	log.Printf("added: %v", payload)

	if len(d.batch) >= d.cfg.MaxBatchSize { // HL
		t := time.Now()
		d.processor.Process(d.batch) // HL
		log.Printf("flushed %d payloads in %s", len(d.batch), time.Since(t))
		d.batch = nil // HL
	}

	return nil
}

// STOP2 OMIT
