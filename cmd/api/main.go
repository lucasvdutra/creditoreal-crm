package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"creditoreal-crm/internal/config"
	"creditoreal-crm/internal/health"
	apphttp "creditoreal-crm/internal/http/middleware"
	"creditoreal-crm/internal/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health.Handle)

	handler := apphttp.RequestID(mux)
	server := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("api.started", slog.String("addr", cfg.HTTPAddr), slog.String("env", cfg.AppEnv))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("api.failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("api.shutdown_failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("api.stopped")
}
