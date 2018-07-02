package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPasswordHandler_Success(t *testing.T) {

	data := url.Values{}
	data.Set("password", "angryMonkey")
	req, err := http.NewRequest("POST", "/hash", strings.NewReader(data.Encode()))
}
