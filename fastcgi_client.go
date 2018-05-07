package main

import (
	"log"

	"github.com/tomasen/fcgi_client"
)

type FastCgiClient struct {
}

func (clinet *FastCgiClient) New(network, address string) *fcgiclient.FCGIClient {
	fc, err := fcgiclient.Dial(network, address)
	if err != nil {
		log.Print(err)
	}
	return fc
}
