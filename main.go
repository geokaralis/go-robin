package main

import (
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
) 

const timeout = 100 * time.Millisecond

func main() {
	server_nl := &Server {
		id: Rand(8),
		url: &url.URL{Host: "127.0.0.1"},
		weight: 20,
	}

	server_us := &Server {
		id: Rand(8),
		url: &url.URL{Host: "127.0.0.1"},
		weight: 40,
	}

	server_nl_handler := NewHandler(server_nl)
	server_us_handler := NewHandler(server_us)

	handler := NewRobin(server_us_handler, server_nl_handler)

	server := &http.Server {
		Handler: handler,
		Addr: ":9090",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Starting web server...")

	req, err := http.NewRequest("GET", "http://localhost:9090/", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Close = true

	var num int
	fmt.Print("How many requests? ")
	_, err = fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatal(err)
	}

	requests := num

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