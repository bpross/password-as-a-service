package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/bpross/password-as-a-service/stats"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStatsHandlerInitial(t *testing.T) {
	req := httptest.NewRequest("GET", "/stats", nil)

	rr := httptest.NewRecorder()
	st := stats.New()
	StatsHandler(rr, req, st)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Total":0,"Average":0}`
	actual := rr.Body.String()
	if actual != expected {
		t.Fatalf("%v", actual)
	}
}

func TestStatsHandlerOneRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/stats", nil)

	rr := httptest.NewRecorder()
	st := stats.New()
	st = addStats(st, 1)
	StatsHandler(rr, req, st)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Total":1,"Average":0}`
	actual := rr.Body.String()
	if actual != expected {
		t.Fatalf("%v", actual)
	}

}

func TestStatsHandlerMultipleRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/stats", nil)

	rr := httptest.NewRecorder()
	st := stats.New()
	numberRequests := 3
	st = addStats(st, numberRequests)
	StatsHandler(rr, req, st)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	readBuf, _ := ioutil.ReadAll(rr.Body)
	r := bytes.NewReader(readBuf)
	decoder := json.NewDecoder(r)
	resp := &StatsResponse{}
	err := decoder.Decode(resp)
	if err != nil {
		t.Fatalf("Failed to decode json")
	}

	if resp.Total != int64(numberRequests) {
		t.Fatalf("Incorrect number of requests %v", resp.Total)
	}
	expectedAverageResponseTime := int64(1000000)
	if resp.Average < expectedAverageResponseTime {
		t.Fatalf("Average Response Time too fast %v", resp.Average)
	}
}

func addStats(st *stats.Stats, numReq int) *stats.Stats {
	for i := 0; i < numReq; i++ {
		start := time.Now()
		wait := time.Duration(i)
		time.Sleep(wait * time.Second)
		st.End(start)
	}
	return st
}
