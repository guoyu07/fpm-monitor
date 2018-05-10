package main

import (
	"io/ioutil"
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

type Monitor struct {
	Connection *fcgiclient.FCGIClient
}

func (monitor *Monitor) NewConnection(netType, address string) *Monitor {
	monitor.Connection = (&FastCgiClient{}).New(netType, address)
	return monitor
}

func (monitor *Monitor) GetFPMStatus(name string, reqParmas map[string]RequestParam) []byte {
	resp, err := monitor.Connection.Get(reqParmas[name])
	defer monitor.Connection.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(string(content))
	defer resp.Body.Close()
	return content
}
