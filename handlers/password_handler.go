package handlers

import (
	pw "github.com/bpross/password-as-a-service/password"
	"net/http"
	"time"
)

const PasswordRequestKey = "password"
const PasswordMissingSlug = "Password is required"

func PasswordHandler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(5 * time.Second)

	plainTextPassword := r.PostFormValue(PasswordRequestKey)
	if plainTextPassword == "" {
		http.Error(w, PasswordMissingSlug, http.StatusBadRequest)
		return
	}

	p := pw.CreateAndHash(plainTextPassword)
	encodedPassword := p.UrlEncode()

	w.Write([]byte(encodedPassword))
}
