package main

import (
	"net/url"
)

const (
	MaxClients = 2048
	ClientTracking = 0 // How long a client will be tracked
) 

type Client struct {
	id string
	server url.URL
	address string
	connections uint64
}

// TODO: Maintain client connections with Robin never surpassing max connection number
func (c *Client) Maintain() {

}
