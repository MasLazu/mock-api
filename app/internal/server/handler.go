package server

import (
	_ "embed"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//go:embed large-response.json
var largeResponse string

func writeResponseMessage(w http.ResponseWriter, status int, message string, handlerCounter prometheus.Gauge) {
	writeResponse(w, status, fmt.Sprintf("{\"message\": \"%s\"}", message), handlerCounter)
}

func writeResponseBody(w http.ResponseWriter, status int, jsonString string, handlerCounter prometheus.Gauge) {
	writeResponse(w, status, fmt.Sprintf("{\"data\": %s}", jsonString), handlerCounter)
}

func writeResponse(w http.ResponseWriter, status int, response string, handlerCounter prometheus.Gauge) {
	requestCounter.Inc()
	if handlerCounter != nil {
		handlerCounter.Inc()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprint(w, response)
}

func delay(r *http.Request) (int, error) {
	seconds, err := strconv.Atoi(r.PathValue("seconds"))
	if err != nil {
		return 0, err
	}

	time.Sleep(time.Duration(seconds) * time.Second)
	return seconds, nil
}

func resetCountersHandler(w http.ResponseWriter, r *http.Request) {
	requestCounter.Set(0)
	getDelayCounter.Set(0)
	getTimeoutCounter.Set(0)
	getSuccessCounter.Set(0)
	getErrorCounter.Set(0)
	getLargeResponseCounter.Set(0)
	postDelayCounter.Set(0)
	postBadrequestDelayCounter.Set(0)

	writeResponseMessage(w, http.StatusOK, "Counters reseted", nil)
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	if seconds, err := delay(r); err != nil {
		writeResponseMessage(w, http.StatusBadRequest, "Invalid seconds parameter", getDelayCounter)
	} else {
		writeResponseMessage(w, http.StatusOK, fmt.Sprintf("Delayed for %d seconds", seconds), getDelayCounter)
	}
}

func timeoutHandler(w http.ResponseWriter, _ *http.Request) {
	requestCounter.Inc()
	getTimeoutCounter.Inc()

	time.Sleep(10 * time.Second)
	w.WriteHeader(http.StatusRequestTimeout)
}

func successHandler(w http.ResponseWriter, _ *http.Request) {
	writeResponseMessage(w, http.StatusOK, "Success!", getSuccessCounter)
}

func errorHandler(w http.ResponseWriter, _ *http.Request) {
	writeResponseMessage(w, http.StatusInternalServerError, "Internal Server Error", getErrorCounter)
}

func postBadRequestDelayHandler(w http.ResponseWriter, r *http.Request) {
	if seconds, err := delay(r); err != nil {
		writeResponseMessage(w, http.StatusBadRequest, "Invalid seconds parameter", postBadrequestDelayCounter)
	} else {
		writeResponseMessage(w, http.StatusBadRequest, fmt.Sprintf("Delayed for %d seconds", seconds), postBadrequestDelayCounter)
	}
}

func postDelayHandler(w http.ResponseWriter, r *http.Request) {
	if seconds, err := delay(r); err != nil {
		writeResponseMessage(w, http.StatusBadRequest, "Invalid seconds parameter", postDelayCounter)
	} else {
		writeResponseMessage(w, http.StatusOK, fmt.Sprintf("Delayed for %d seconds", seconds), postDelayCounter)
	}
}

func largeResponseHandler(w http.ResponseWriter, r *http.Request) {
	writeResponseBody(w, http.StatusOK, largeResponse, getLargeResponseCounter)
}
