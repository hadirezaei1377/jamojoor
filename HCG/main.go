package hcg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// prometheus metrics

var (
	// Counter: number of ping call
	pingCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ping_requests_total",
			Help: "Total number of ping requests received.",
		},
	)

	// Gauge: status of active requests and number of that
	activeRequestsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_requests",
			Help: "Current number of active requests being processed.",
		},
	)

	// Histogram: measure the request processing time
	requestDurationHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Histogram of request durations for /ping endpoint.",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(pingCounter)
	prometheus.MustRegister(activeRequestsGauge)
	prometheus.MustRegister(requestDurationHistogram)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()

	activeRequestsGauge.Inc()
	defer activeRequestsGauge.Dec()

	timer := prometheus.NewTimer(requestDurationHistogram)
	defer timer.ObserveDuration()

	fmt.Fprintln(w, "pong")
}

func main() {

	http.HandleFunc("/ping", pingHandler)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
