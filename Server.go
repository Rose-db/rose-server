package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	Host string
	Port string
}

func (s *Server) Start() {
	http.HandleFunc("/", s.Handle)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), nil)

	if err != nil {
		fmt.Printf("An error occurred when starting Rose: %s\n", err.Error())

		return
	}
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.HandleError(w, r)
		return
	}
}

func (s *Server) HandleError(w http.ResponseWriter, r *http.Request) {
}
