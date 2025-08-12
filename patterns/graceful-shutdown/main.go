package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			log.Fatalf("failed to write to the server. Reason: %v", err)
		}
	})

	server := &http.Server{
		Addr:    ":9000",
		Handler: mux,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Could not listen. Reason: %v", err)
		}
	}()

	sig := <-quit
	log.Printf("Received signal: %v. Shutting down server...\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	log.Fatal(server.Shutdown(ctx))
}
