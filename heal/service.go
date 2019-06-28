package heal

import (
	"fmt"
	"net/url"
	"net/http"
)

type State int

const (
	Running State = iota
	Unavailable
)

type Service struct {
	Id string
	Url *url.URL
	Endpoint string
	State int
	Port string
}

func (s *Service) Start() {
	go func() {
		http.HandleFunc(s.Endpoint, func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Healing service %d, %s", s.State, s.Id)
		})
		http.ListenAndServe(s.Port, nil)
	}()
}

