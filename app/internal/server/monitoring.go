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
		metrics := initMetrics()
		prometheus.MustRegister(metrics...)

		http.Handle("/metrics", promhttp.Handler())
		slog.Info("Starting exporter on :9700")
		log.Fatal(http.ListenAndServe(":9700", nil))
	})
}
