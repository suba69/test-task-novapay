package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"test-task/http"
)

func main() {
	router := http.NewRouter()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Starting server on port 8080...")
		if err := http.Server(ctx, router); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
}
