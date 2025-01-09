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

var (
	requestCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "get_request_counter",
		Help: "A counter for the total number of requests",
	})

	getDelayCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "get_delay_counter",
		Help: "A counter for the get delay endpoint",
	})

	getTimeoutCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "get_timeout_counter",
		Help: "A counter for the get timeout endpoint",
	})

	getSuccessCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "get_success_counter",
		Help: "A counter for the get success endpoint",
	})

	getErrorCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "get_error_counter",
		Help: "A counter for the get error endpoint",
	})

	getLargeResponseCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "get_largeres_counter",
		Help: "A counter for the get large-response endpoint",
	})

	postDelayCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "post_delay_counter",
		Help: "A counter for the post delay endpoint",
	})

	postBadrequestDelayCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "post_badreq_delay_counter",
		Help: "A counter for the post delay endpoint",
	})
)

func monitorMetrics() {
	once.Do(func() {
		prometheus.MustRegister(
			requestCounter,
			getDelayCounter,
			getTimeoutCounter,
			getSuccessCounter,
			getErrorCounter,
			getLargeResponseCounter,
			postDelayCounter,
			postBadrequestDelayCounter,
		)

		http.Handle("/metrics", promhttp.Handler())
		slog.Info("Starting exporter on :9700")
		log.Fatal(http.ListenAndServe(":9700", nil))
	})
}
