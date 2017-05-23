package dispatcher

import (
	"net/http"
	"encoding/json"
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


