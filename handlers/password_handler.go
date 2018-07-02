package handlers

import (
	pw "github.com/bpross/password-as-a-service/password"
	"net/http"
	"os"
	"strconv"
	"time"
)

const PasswordWait = "PASSWORDWAIT"
const FallbackWait = 5
const PasswordRequestKey = "password"
const PasswordMissingSlug = "Password is required"

func PasswordHandler(w http.ResponseWriter, r *http.Request) {

	defaultSleepTime := getSleepTime(PasswordWait, FallbackWait)
	time.Sleep(defaultSleepTime * time.Second)

	plainTextPassword := r.PostFormValue(PasswordRequestKey)
	if plainTextPassword == "" {
		http.Error(w, PasswordMissingSlug, http.StatusBadRequest)
		return
	}

	p := pw.CreateAndHash(plainTextPassword)
	encodedPassword := p.UrlEncode()

	w.Write([]byte(encodedPassword))
}

func getSleepTime(key string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		valueInt, _ := strconv.Atoi(value)
		return time.Duration(valueInt)
	}
	return fallback
}
