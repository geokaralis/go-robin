package health

import (
	"fmt"
	"net/http"
)

const (
	Endpoint = "/health"
	Port = ":3001"
)

type Health struct {
	Instance string
	Endpoint string
	Port string
}

func (h *Health) ServeHealth()  {
	go func() {
		http.HandleFunc(h.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Health endpoint, instance: [%s]", h.Instance)
		})
		http.ListenAndServe(h.Port, nil)
	}()

	fmt.Println("Serving health...")
}

func NewHealth() *Health {
	return &Health {
		Instance: "",
		Endpoint: Endpoint,
		Port: Port,
	}
}