package dispatcher

import (
	"encoding/json"
	"net/http"
)

func NewEventHandler(d *Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var event event

		if err := json.NewDecoder(req.Body).Decode(&event); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		d.Add(event)
	}
}
