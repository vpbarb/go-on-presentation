package server

import (
	"encoding/json"
	"net/http"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher/dispatcher.v4"
)

// START1 OMIT
func NewDispatcherHandler(d *dispatcher.Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var p dispatcher.Payload

		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := d.Collect(p); err != nil { // HL
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}
}

// STOP1 OMIT
