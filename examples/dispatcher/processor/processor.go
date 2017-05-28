package processor

import (
	"math/rand"
	"time"
)

// START1 OMIT
type (
	Item  map[string]string
	Batch []Item
	Fake  struct{}
)

func (f *Fake) Process(batch Batch) error {
	if len(batch) == 0 {
		return nil
	}
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Fake process delay
	return nil
}

// STOP1 OMIT
