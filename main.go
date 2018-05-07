package main

import (
	"io/ioutil"
	"log"
)

func main() {
	client := &FastCgiClient{}
	fc := client.New("tcp", "127.0.0.1:9000")
	requestParam := make(map[string]string)
	requestParam["REQUEST_METHOD"] = "GET"
	requestParam["SCRIPT_NAME"] = "/status"
	requestParam["SCRIPT_FILENAME"] = "/status"
	requestParam["QUERY_STRING"] = "json&full"
	resp, err := fc.Get(requestParam)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("content:", string(content))
	defer resp.Body.Close()
}
