package main

import (
	"context"
	"github.com/bpross/password-as-a-service/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	httpServer, logger := setup()

	go func() {
		logger.Printf("Listening on http://0.0.0.0%s\n", httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()
	graceful(httpServer, logger)
}

func setup() (*http.Server, *log.Logger) {
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":8080"
	}

	logger := log.New(os.Stdout, "", 0)

	return &http.Server{
		Addr:    addr,
		Handler: server.New(server.Logger(logger)),
	}, logger
}

func graceful(hs *http.Server, logger *log.Logger) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Printf("\nShutdown!\n")

	if err := hs.Shutdown(ctx); err != nil {
		logger.Printf("Error: %v\n", err)
	} else {
		logger.Println("Server stopped")
	}
}
