package processor

import ( // OMIT
	"math/rand" // OMIT
	"time"      // OMIT
) // OMIT
// OMIT
type (
	Batch [][]byte
	Fake  struct{}
)

func (f *Fake) Process(batch Batch) {
	if len(batch) == 0 {
		return
	}
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Fake process time
}
