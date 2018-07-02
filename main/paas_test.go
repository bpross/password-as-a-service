package main

import (
	"github.com/bpross/password-as-a-service/handlers"
	"syscall"
	"testing"
	"time"
)

func TestGraceful(t *testing.T) {
	go func() {
		time.Sleep(10 * time.Millisecond)
		handlers.ShutdownChannel <- syscall.SIGTERM
	}()

	main()
}
