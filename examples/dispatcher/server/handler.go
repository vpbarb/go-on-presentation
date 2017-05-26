package server

import (
	"encoding/json"
	"net/http"

	"github.com/Barberrrry/go-on-presentation/examples/dispatcher"
)

func NewHandler(d *dispatcher.Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var p dispatcher.Payload

		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		d.Add(p)
	}
}
