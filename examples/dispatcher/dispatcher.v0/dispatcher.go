package dispatcher

import (
	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/processor"
)

type (
	// Any processor must implement this interface. Fake processor does.
	Processor interface {
		Process(processor.Batch)
	}

	// Type of some key-value data received by HTTP requests in JSON format
	Payload map[string]string
)
