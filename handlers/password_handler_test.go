package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPasswordHandlerSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "/password", nil)
	if err != nil {
		t.Fatal(err)
	}
	data := url.Values{}
	data.Add("password", "angryMonkey")
	req.PostForm = data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PasswordHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	actual := rr.Body.String()
	if actual != expected {
		t.Fatalf("%s != %s", actual, expected)
	}
}

func TestPasswordHandlerEmptyPassword(t *testing.T) {
	req, err := http.NewRequest("POST", "/password", nil)
	if err != nil {
		t.Fatal(err)
	}
	data := url.Values{}
	data.Add("password", "")
	req.PostForm = data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PasswordHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPasswordHandlerNoPassword(t *testing.T) {
	req, err := http.NewRequest("POST", "/password", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PasswordHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
