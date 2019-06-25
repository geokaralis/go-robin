package main

import (
	"net/url"
)

type Server struct {
	id string
	url *url.URL
	weight int
}

func NewServer(url *url.URL) *Server {
	return &Server {
		id: "",
		url: url,
		weight: 0,
	}
}