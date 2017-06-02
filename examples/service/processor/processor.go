package processor

import ( // OMIT
	"math/rand" // OMIT
	"time"      // OMIT
) // OMIT
// OMIT
type Batch []string // HL

type Fake struct{} // HL

func (f *Fake) Process(batch Batch) {
	if len(batch) == 0 {
		return
	}
	// Fake process time
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // HL
}
