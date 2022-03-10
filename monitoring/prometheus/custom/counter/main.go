package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "simple_processed_opts_total",
		Help: "The total number of processed events",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

	// curl localhost:2112/metrics
	//# HELP simple_processed_opts_total The total number of processed events
	//# TYPE simple_processed_opts_total counter
	//simple_processed_opts_total 12
}
