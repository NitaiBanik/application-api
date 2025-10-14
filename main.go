package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"application-api/config"
	"application-api/handlers"
	"application-api/middleware"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	log.Printf("Starting application-api on port %s", cfg.Port)

	handler := handlers.NewHandler()

	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware())

	router.HandleFunc("/testapi", handler.TestApiHandler).Methods("POST")
	router.HandleFunc("/health", handler.HealthHandler).Methods("GET")

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
