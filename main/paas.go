package main

import (
	h "github.com/bpross/password-as-a-service/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/hash", h.PasswordHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
