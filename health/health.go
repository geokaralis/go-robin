package health

import (
	"fmt"
	"net/http"
)

const (
	endpoint = "/health"
	port = ":3001"
)

type Health struct {
	instance string
	endpoint string
	port string
}

func (h *Health) ServeHealth()  {
	http.HandleFunc(h.endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health endpoint, instance: [%s]", h.instance)
	})
	http.ListenAndServe(h.port, nil)
}

func NewHealth() *Health {
	return &Health {
		instance: "",
		endpoint: endpoint,
		port: port,
	}
}