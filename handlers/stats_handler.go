package handlers

import (
	"encoding/json"
	"github.com/bpross/password-as-a-service/stats"
	"net/http"
)

type StatsResponse struct {
	Total   int64
	Average int64
}

func StatsHandler(w http.ResponseWriter, r *http.Request, st *stats.Stats) {
	totalResponses := st.PasswordReponses
	averageResponseTime := st.AverageResponseTime()
	statsResponse := StatsResponse{totalResponses, averageResponseTime}

	js, err := json.Marshal(statsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
