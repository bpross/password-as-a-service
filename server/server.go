package server

import (
	h "github.com/bpross/password-as-a-service/handlers"
	"github.com/bpross/password-as-a-service/stats"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
	stats  *stats.Stats
}

func New(options ...func(*Server)) *Server {
	s := &Server{
		logger: log.New(os.Stdout, "", 0),
		mux:    http.NewServeMux(),
		stats:  stats.New(),
	}

	for _, f := range options {
		f(s)
	}

	s.mux.HandleFunc("/hash", s.HashHandler)
	s.mux.HandleFunc("/shutdown", s.ShutdownHandler)
	s.mux.HandleFunc("/stats", s.StatsHandler)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func Logger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

func (s *Server) HashHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	s.logger.Print("Url Encoding password")
	h.PasswordHandler(w, r)
	s.stats.End(start)
}

func (s *Server) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Print("Received Shutdown Request")
	h.ShutdownHandler(w, r)
}

func (s *Server) StatsHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Print("Received Stats Request")
	h.StatsHandler(w, r, s.stats)
}
