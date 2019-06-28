package main

import (
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"go-robin/health"
	"go-robin/queue"
	"go-robin/heal"
) 

const timeout = 100 * time.Millisecond

func main() {
	fmt.Println("Press 1 to run")
	fmt.Println("Press 2 to exit")
	setup()
	for {
			sample()
	}
}

func sample() {
	var input string
	fmt.Printf("Input (req, exit): ")
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
			 fmt.Println("invalid input")
			 return
	}
	switch input {
	case "req":
			request()
	case "exit":
			os.Exit(2)
	default:
			fmt.Println("def")
	}
}

func setup() {

	healService := &heal.Service {
		Id: Rand(4),
		Url: &url.URL{Host: "127.0.0.1"},
		Endpoint: "/heal",
		State: 1,
		Port: ":1251",
	}

	healService.Start()

	heal.StartHealingProcess()


	messageQueue := queue.Queue{}
	messageQueue.NewQueue()

	messageQueue.Enqueue("Hello")
	messageQueue.Enqueue("World")
	messageQueue.Enqueue("!")

	fmt.Printf("Queue size: %d\n", messageQueue.Size())

	fmt.Printf("Queue's first item: %v\n", *messageQueue.Front())

	for i := 0; i < messageQueue.Size(); i++ {
		fmt.Printf("%v ", messageQueue.Items[i])
	}
	fmt.Println()
	fmt.Println()

	session := &Session {
		id: Rand(6),
		status: 1,
		pid: os.Getpid(),
		name: "Robin",
		desc: "Decentralized load balancer in Go.",
	}

	fmt.Printf("Process: %d, Name: %s, Description: %s\n", session.pid, session.name, session.desc)

	session.Watch()

	WatchProcess()

	health_server := &health.Health {
		Instance: Rand(4),
		Endpoint: "/health",
		Port: ":3001",
	}

	health_server.ServeHealth()

	connection := &Connection {
		id: Rand(8),
		Type: "TCP",
		status: 1,
	}

	connection.Check()
	
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

}

func request() {
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