package main

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Handler:           handler,
		Addr:              ":" + port,
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 10,
	}
	return s.httpServer.ListenAndServe()
}
