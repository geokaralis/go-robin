package main

import (
	"fmt"
	"net/http"
	"sync"
) 

type Robin struct {
	current int
	mu *sync.Mutex
	handlers []http.Handler
}

func (r *Robin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mu.Lock()

	defer r.mu.Unlock()

	r.handlers[r.current].ServeHTTP(w, req)

	r.current = (r.current + 1) % len(r.handlers)
}

func (r *Robin) AddHandler(handler http.Handler) {
	r.mu.Lock()

	defer r.mu.Unlock()

	r.handlers = append(r.handlers, handler)
}

func NewRobin(handlers ...http.Handler) *Robin {
	return &Robin{
		current: 0,
		mu: &sync.Mutex{},
		handlers: handlers,
	}
}

func NewHandler(server *Server) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, from instance:[%s], at host:[%s]", server.id, server.url.Host)
	})
	return mux
}