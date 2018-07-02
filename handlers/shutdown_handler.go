package handlers

import (
	"net/http"
	"os"
	"syscall"
)

var ShutdownChannel = make(chan os.Signal, 1)

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	ShutdownChannel <- syscall.SIGTERM
}
