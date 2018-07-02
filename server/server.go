package server

import (
	h "github.com/bpross/password-as-a-service/handlers"
	"log"
	"net/http"
)

type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}

// Used this https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2#file-graceful-go-L17
func NewServer(options ...func(*Server)) *Server {
	s := &Server{
		logger: log.New(os.Stdout, "", 0),
		mux:    http.NewServeMux(),
	}

	for _, f := range options {
		f(s)
	}

	s.mux.HandleFunc("/hash", h.PasswordHandler)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
