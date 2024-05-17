package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "myapp_requests_total",
            Help: "Total number of requests received.",
        },
        []string{"path"},
    )
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "myapp_request_duration_seconds",
            Help: "Duration of HTTP requests.",
        },
        []string{"path"},
    )
)

func main() {
    prometheus.MustRegister(requests)
    prometheus.MustRegister(requestDuration)

    http.HandleFunc("/", handleRequest)

	// Serve the metrics endpoint on the "/metrics" path
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    // Simulate variable load
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    w.Write([]byte("Hello, World!"))
    duration := time.Since(start)

    requests.WithLabelValues(r.URL.Path).Inc()
    requestDuration.WithLabelValues(r.URL.Path).Observe(duration.Seconds())
}