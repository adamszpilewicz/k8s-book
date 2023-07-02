package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	appPort = 8888
)

func main() {
	for {
		url := fmt.Sprintf("http://localhost:%d", appPort)
		resp, err := http.Get(fmt.Sprintf("%s/status", url))
		if err != nil {
			log.Printf("Could not reach the service: %s\n", err.Error())
			log.Printf("Service Status: FAIL\n")
			time.Sleep(5 * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Received status code %d\n", resp.StatusCode)
			log.Printf("Service Status: FAIL\n")
		} else {
			log.Printf("Service Status: OK\n")
		}

		time.Sleep(5 * time.Second) // wait for 5 seconds before the next check
	}
}
