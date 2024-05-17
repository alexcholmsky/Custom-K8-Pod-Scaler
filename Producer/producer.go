package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
)

var (
    endpoint   = flag.String("endpoint", "http://my-go-app:8080", "Endpoint to send requests to")
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
        } else if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			log.Printf("Successfully reached endpoint: %v, Status Code: %d\n", *endpoint, resp.StatusCode)
		
			// Read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading response body: %v\n", err)
			} else {
				log.Printf("Response Body: %s\n", body)
			}
		
			// Ensure body is closed after handling the response
			resp.Body.Close()
			consecutiveFailures = 0  // Reset the failure count on a successful request
		} else {
			log.Printf("Request to endpoint %v failed with status code: %d\n", *endpoint, resp.StatusCode)
			resp.Body.Close()
			consecutiveFailures++  // Increase the failure count on a failed request
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

