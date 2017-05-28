package dispatcher

import (
	"log"
	"sync"
	"time"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

// START1 OMIT
type (
	Processor interface {
		Process(processor.Batch) error
	}
	Payload    map[string]string
	Dispatcher struct {
		processor    Processor
		maxBatchSize int
		flushLock    sync.Mutex
		batch        processor.Batch
	}
	Config struct {
		MaxBatchSize int
	}
)

// STOP1 OMIT

// START2 OMIT
func New(cfg Config) *Dispatcher {
	return &Dispatcher{
		processor:    &processor.Fake{},
		maxBatchSize: cfg.MaxBatchSize,
	}
}

// STOP2 OMIT

// START3 OMIT
func (d *Dispatcher) Add(payload Payload) error {
	log.Printf("add payload: %v", payload)

	d.flushLock.Lock()
	defer d.flushLock.Unlock()

	d.batch = append(d.batch, processor.Item(payload))
	if len(d.batch) >= d.maxBatchSize {
		t := time.Now()
		d.processor.Process(d.batch)
		log.Printf("proceed %d payloads in %s", len(d.batch), time.Since(t))
		d.batch = nil
	}

	return nil
}

// STOP3 OMIT
