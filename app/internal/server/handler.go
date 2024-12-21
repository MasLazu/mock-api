package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "request_counter",
		Help: "A counter for the total number of requests",
	})

	delayCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "delay_counter",
		Help: "A counter for the delay endpoint",
	})

	timeoutCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "timeout_counter",
		Help: "A counter for the timeout endpoint",
	})

	successCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "success_counter",
		Help: "A counter for the success endpoint",
	})

	errorCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "error_counter",
		Help: "A counter for the error endpoint",
	})
)

func writeResponseMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "{\"message\": \"%s\"}", message)
}

func resetCountersHandler(w http.ResponseWriter, r *http.Request) {
	requestCounter.Set(0)
	delayCounter.Set(0)
	timeoutCounter.Set(0)
	successCounter.Set(0)
	errorCounter.Set(0)

	writeResponseMessage(w, http.StatusOK, "Counters reseted")
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	requestCounter.Inc()
	delayCounter.Inc()

	seconds, err := strconv.Atoi(r.PathValue("seconds"))
	if err != nil {
		writeResponseMessage(w, http.StatusBadRequest, "Invalid seconds parameter")
		return
	}
	time.Sleep(time.Duration(seconds) * time.Second)

	writeResponseMessage(w, http.StatusOK, fmt.Sprintf("Delayed for %d seconds", seconds))
}

func timeoutHandler(w http.ResponseWriter, _ *http.Request) {
	requestCounter.Inc()
	timeoutCounter.Inc()

	time.Sleep(10 * time.Second)
	w.WriteHeader(http.StatusRequestTimeout)
}

func successHandler(w http.ResponseWriter, _ *http.Request) {
	requestCounter.Inc()
	successCounter.Inc()

	writeResponseMessage(w, http.StatusOK, "Success!")
}

func errorHandler(w http.ResponseWriter, _ *http.Request) {
	requestCounter.Inc()
	errorCounter.Inc()

	writeResponseMessage(w, http.StatusInternalServerError, "Internal Server Error")
}
