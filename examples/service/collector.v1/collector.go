// START1 OMIT
package collector

import ( // OMIT
	"log"  // OMIT
	"time" // OMIT
	// OMIT
	"github.com/Barberrrry/go-on-presentation/examples/service/processor" // OMIT
) // OMIT
// OMIT
type (
	Processor interface {
		Process(processor.Batch)
	}
	Collector struct {
		cfg       Config
		processor Processor
		batch     processor.Batch
	}
	Config struct {
		MaxBatchSize int
	}
)

func New(cfg Config, processor Processor) *Collector {
	return &Collector{cfg: cfg, processor: processor}
}

// STOP1 OMIT

// START2 OMIT
func (c *Collector) Collect(payload []byte) error {
	c.batch = append(c.batch, payload)
	log.Printf("collected: %s", payload)

	if len(c.batch) >= c.cfg.MaxBatchSize { // HL
		t := time.Now()
		c.processor.Process(c.batch) // HL
		log.Printf("flushed %d payloads in %s", len(c.batch), time.Since(t))
		c.batch = nil // HL
	}

	return nil
}

// STOP2 OMIT
