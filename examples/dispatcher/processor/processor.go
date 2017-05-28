package processor

import (
	"math/rand"
	"time"
)

type (
	Item  map[string]string
	Batch []Item
	Fake  struct{}
)

func (f *Fake) Process(batch Batch) {
	if len(batch) == 0 {
		return
	}
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Fake process time
}
