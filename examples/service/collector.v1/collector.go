// START1 OMIT
package collector

import ( // OMIT
	"log"  // OMIT
	"time" // OMIT
	// OMIT
	"github.com/Barberrrry/go-on-presentation/examples/service/processor" // OMIT
) // OMIT
// OMIT
type Processor interface {
	Process(processor.Batch)
}

type Collector struct {
	Processor    Processor
	MaxBatchSize int
	buffer       processor.Batch
}

// STOP1 OMIT

// START2 OMIT
func (c *Collector) Collect(payload string) {
	c.buffer = append(c.buffer, payload) // HL

	log.Printf("collected: %s", payload)

	if len(c.buffer) >= c.MaxBatchSize { // HL
		c.flush(c.buffer) // HL
		c.buffer = nil    // HL
	} // HL
}

func (c *Collector) flush(batch processor.Batch) {
	t := time.Now()
	c.Processor.Process(batch) // HL
	log.Printf("flushed %d payloads in %s", len(batch), time.Since(t))
}

// STOP2 OMIT
