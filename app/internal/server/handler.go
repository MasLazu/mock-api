package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "request_counter",
		Help: "A counter for the total number of requests",
	})

	delayCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "delay_counter",
		Help: "A counter for the delay endpoint",
	})

	timeoutCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "timeout_counter",
		Help: "A counter for the timeout endpoint",
	})

	successCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "success_counter",
		Help: "A counter for the success endpoint",
	})

	errorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "error_counter",
		Help: "A counter for the error endpoint",
	})
)

func writeResponseMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "{\"message\": \"%s\"}", message)
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
	time.Sleep(30 * time.Second)
	writeResponseMessage(w, http.StatusOK, "This probably won't be seen due to timeout")
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
