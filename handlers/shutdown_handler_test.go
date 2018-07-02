package handlers

import (
	"net/http"
	"net/http/httptest"
	"syscall"
	"testing"
)

func TestShutdownHandler(t *testing.T) {
	req, err := http.NewRequest("PUT", "/shutdown", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ShutdownHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	s := <-ShutdownChannel
	if s != syscall.SIGTERM {
		t.Fatalf("Shtudown Handler did not send SIGTERM")
	}
}
