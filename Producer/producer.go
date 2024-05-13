package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
    endpoint   = flag.String("endpoint", "http://example.com", "Endpoint to send requests to")
    mode       = flag.String("mode", "http", "Operating mode: 'http' or 'queue'")
    rate       = flag.Int("rate", 10, "Requests per second")
)

func httpRequests() {
    flag.Parse()
    ticker := time.NewTicker(time.Second / time.Duration(*rate))
    defer ticker.Stop()

    consecutiveFailures := 0  // Counter for consecutive failures

    for range ticker.C {
        resp, err := http.Get(*endpoint)
        if err != nil {
            log.Printf("Failed to reach endpoint: %v\n", err)
            consecutiveFailures++
            if consecutiveFailures >= 10 {
                log.Fatalf("Failed to reach endpoint 10 times in a row: %v\n", err)
            }
        } else {
            log.Printf("Successfully reached endpoint: %v, Status Code: %d\n", *endpoint, resp.StatusCode)
            // Ensure body is closed after handling the response
            resp.Body.Close()
            consecutiveFailures = 0  // Reset the failure count on a successful request
        }
    }
}

func messageQueue(){
	fmt.Println("Not implemented")
}

func main() {
	if *mode == "http" {
		httpRequests()
	} else {
		messageQueue()
	}
}

