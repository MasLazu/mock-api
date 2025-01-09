package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func Run() error {
	setupLogger()

	mux := http.NewServeMux()
	setupRoutes(mux)

	go monitorMetrics()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", getPort()),
		Handler: mux,
	}

	slog.Info("Starting server", "addr", server.Addr)
	return server.ListenAndServe()
}

func setupLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /delay/{seconds}", delayHandler)
	mux.HandleFunc("GET /timeout", timeoutHandler)
	mux.HandleFunc("GET /success", successHandler)
	mux.HandleFunc("GET /error", errorHandler)
	mux.HandleFunc("GET /reset", resetCountersHandler)

	mux.HandleFunc("POST /delay/{seconds}", delayHandler)
	mux.HandleFunc("POST /badrequest/delay/{seconds}", badRequestDelayHandler)
	mux.HandleFunc("POST /timeout", timeoutHandler)
	mux.HandleFunc("POST /success", successHandler)
	mux.HandleFunc("POST /error", errorHandler)
	mux.HandleFunc("POST /reset", resetCountersHandler)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "80"
	}
	return port
}
