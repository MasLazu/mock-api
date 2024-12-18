package server

import (
	"log"
	"log/slog"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var once sync.Once

func monitorMetrics() {
	once.Do(func() {
		prometheus.MustRegister(
			requestCounter,
			delayCounter,
			timeoutCounter,
			successCounter,
			errorCounter,
		)

		http.Handle("/metrics", promhttp.Handler())
		slog.Info("Starting exporter on :9700")
		log.Fatal(http.ListenAndServe(":9700", nil))
	})
}
