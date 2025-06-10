// Package main is an entry point to application
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/artnikel/iotask/internal/api"
	"github.com/artnikel/iotask/internal/config"
	"github.com/artnikel/iotask/internal/constants"
	"github.com/artnikel/iotask/internal/logging"
	"github.com/artnikel/iotask/internal/models"
	"github.com/artnikel/iotask/internal/service"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger, err := logging.NewLogger(cfg.Logging.Path)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	manager := &models.Manager{
		Task: make(map[string]*models.Task),
	}
	taskService := service.NewTaskService(manager)
	h := api.NewHandler(taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateTaskHandler(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetTaskHandler(w, r)
		case http.MethodDelete:
			h.DeleteTaskHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  constants.ServerTimeout,
		WriteTimeout: constants.ServerTimeout,
	}

	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), constants.ServerTimeout)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Error.Fatalf("http server shutdown error %v", err)
		}
		close(stopped)
	}()

	log.Printf("starting HTTP server on :%d\n", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		logger.Error.Fatalf("http server not listening: %v", err)
	}

	<-stopped
}
