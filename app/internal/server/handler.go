package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounterChannel = make(chan prometheus.Gauge, 1)
	delayCounterChannel   = make(chan prometheus.Gauge, 1)
	timeoutCounterChannel = make(chan prometheus.Gauge, 1)
	successCounterChannel = make(chan prometheus.Gauge, 1)
	errorCounterChannel   = make(chan prometheus.Gauge, 1)
)

func initMetrics() []prometheus.Collector {
	requestCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "request_counter",
		Help: "A counter for the total number of requests",
	})

	delayCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "delay_counter",
		Help: "A counter for the delay endpoint",
	})

	timeoutCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "timeout_counter",
		Help: "A counter for the timeout endpoint",
	})

	successCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "success_counter",
		Help: "A counter for the success endpoint",
	})

	errorCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "error_counter",
		Help: "A counter for the error endpoint",
	})

	requestCounterChannel <- requestCounter
	delayCounterChannel <- delayCounter
	timeoutCounterChannel <- timeoutCounter
	successCounterChannel <- successCounter
	errorCounterChannel <- errorCounter

	return []prometheus.Collector{requestCounter, delayCounter, timeoutCounter, successCounter, errorCounter}
}

func incrementTotalMetrics() {
	requestCounter := <-requestCounterChannel
	requestCounter.Inc()
	requestCounterChannel <- requestCounter
}

func writeResponseMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "{\"message\": \"%s\"}", message)
}

func resetCountersHandler(w http.ResponseWriter, r *http.Request) {
	requestCounter := <-requestCounterChannel
	requestCounter.Set(0)
	requestCounterChannel <- requestCounter

	delayCounter := <-delayCounterChannel
	delayCounter.Set(0)
	delayCounterChannel <- delayCounter

	timeoutCounter := <-timeoutCounterChannel
	timeoutCounter.Set(0)
	timeoutCounterChannel <- timeoutCounter

	successCounter := <-successCounterChannel
	successCounter.Set(0)
	successCounterChannel <- successCounter

	errorCounter := <-errorCounterChannel
	errorCounter.Set(0)
	errorCounterChannel <- errorCounter

	writeResponseMessage(w, http.StatusOK, "Counters reseted")
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	incrementTotalMetrics()
	delayCounter := <-delayCounterChannel
	delayCounter.Inc()
	delayCounterChannel <- delayCounter

	seconds, err := strconv.Atoi(r.PathValue("seconds"))
	if err != nil {
		writeResponseMessage(w, http.StatusBadRequest, "Invalid seconds parameter")
		return
	}
	time.Sleep(time.Duration(seconds) * time.Second)

	writeResponseMessage(w, http.StatusOK, fmt.Sprintf("Delayed for %d seconds", seconds))
}

func timeoutHandler(w http.ResponseWriter, _ *http.Request) {
	incrementTotalMetrics()
	timeoutCounter := <-timeoutCounterChannel
	timeoutCounter.Inc()
	timeoutCounterChannel <- timeoutCounter

	time.Sleep(10 * time.Second)
	w.WriteHeader(http.StatusRequestTimeout)
}

func successHandler(w http.ResponseWriter, _ *http.Request) {
	incrementTotalMetrics()
	successCounter := <-successCounterChannel
	successCounter.Inc()
	successCounterChannel <- successCounter

	writeResponseMessage(w, http.StatusOK, "Success!")
}

func errorHandler(w http.ResponseWriter, _ *http.Request) {
	incrementTotalMetrics()
	errorCounter := <-errorCounterChannel
	errorCounter.Inc()
	errorCounterChannel <- errorCounter

	writeResponseMessage(w, http.StatusInternalServerError, "Internal Server Error")
}
