package main

import (
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
) 

const timeout = 100 * time.Millisecond

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

func NewHandler(msg string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, from %s", msg)
	})
	return mux
}

func main() {
	server_us := NewHandler("127.0.0.2")
	server_eu := NewHandler("127.0.0.3")
	server_eu_gb := NewHandler("127.0.0.4")

	handler := NewRobin(server_us, server_eu, server_eu_gb)

	server := &http.Server {
		Handler: handler,
		Addr: "127.0.0.1:9090",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Starting web server...")

	req, err := http.NewRequest("GET", "http://127.0.0.1:9090/", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Close = true

	requests := 5

	fmt.Printf("Requests: [%d]\n", requests)

	for i := 0; i < requests; i++ {
		fmt.Printf("Calling... (%d/%d)\n", i+1, requests)
		Lookup(req)
	}

	fmt.Println("Completed all requests")
}

func Lookup(req *http.Request) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received response, %s\n", b)
}