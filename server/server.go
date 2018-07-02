package server

import (
	h "github.com/bpross/password-as-a-service/handlers"
	"log"
	"net/http"
	"os"
)

type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}

// Used this https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2#file-graceful-go-L17
func New(options ...func(*Server)) *Server {
	s := &Server{
		logger: log.New(os.Stdout, "", 0),
		mux:    http.NewServeMux(),
	}

	for _, f := range options {
		f(s)
	}

	s.mux.HandleFunc("/hash", s.HashHandler)
	s.mux.HandleFunc("/shutdown", s.ShutdownHandler)

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
	s.logger.Print("Url Encoding password")
	h.PasswordHandler(w, r)
}

func (s *Server) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Print("Received Shutdown Request")
	h.ShutdownHandler(w, r)
}
