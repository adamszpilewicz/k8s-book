package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://localhost:8081/status")
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
