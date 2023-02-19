package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/monta-k/go-openapi-gen-example/internal/movie"
)

func main() {
	// server
	server, err := movie.NewServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to prepare server, %s", err.Error())
		return
	}
	fmt.Printf("start movie server\n")
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "closed user server, %s", err.Error())
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	fmt.Printf("SIGNAL %d received, so server shutting down now...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("failed to gracefully shutdown, %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "server shutdown completed")
}
