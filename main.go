package main

import (
	"log"
	"time"
)

type RequestParam map[string]string

func main() {
	reqParma := make(map[string]RequestParam)
	requestParam := make(map[string]string)
	requestParam["REQUEST_METHOD"] = "GET"
	requestParam["SCRIPT_NAME"] = "/status"
	requestParam["SCRIPT_FILENAME"] = "/status"
	requestParam["QUERY_STRING"] = "json&full"
	reqParma["localhost"] = requestParam

	monitor := &Monitor{}

	for {
		select {
		case <-time.After(time.Second * 30):
			content := monitor.NewConnection("tcp", "127.0.0.1:9000").GetFPMStatus("localhost", reqParma)
			log.Printf("%s", string(content))
			break
		}
	}
}
