package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ServerConfig struct {
	TLSCertPath string
	TLSKeyPath  string
	Port        int
	TLSEnabled  bool
}

type Server interface {
	ListenAndServe()
}

type server struct {
	handler http.Handler
	config  ServerConfig
}

func NewServer(endpoints []Endpoints, config *ServerConfig) Server {
	if config == nil {
		config = &ServerConfig{}
	}
	if config.Port == 0 {
		config.Port = 8080
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	for _, e := range endpoints {
		e.Register(router)
	}

	return &server{
		config: *config,
		handler: cors.New(
			cors.Options{
				AllowCredentials: true,
				AllowedOrigins:   []string{"mly.vercel.app"},
				AllowedMethods:   []string{"DELETE", "GET", "POST", "PUT"},
				AllowedHeaders:   []string{"Authorization", "Content-Type"},
			},
		).Handler(router),
	}
}

func (s *server) ListenAndServe() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.handler,
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<-sig
		log.Println("Shutting down server...")

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
