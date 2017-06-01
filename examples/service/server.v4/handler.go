package server

import (
	"io/ioutil"
	"net/http"

	"github.com/Barberrrry/go-on-presentation/examples/service/collector.v4"
)

// START1 OMIT
func NewCollectorHandler(d *collector.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := ioutil.ReadAll(req.Body) // HL
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := d.Collect(string(payload)); err != nil { // HL
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}
}

// STOP1 OMIT
