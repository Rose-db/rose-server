package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HttpErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Server struct {
	Host string
	Port string
	Server *http.Server
	App *AppController
}

func (s *Server) Validate(r *http.Request) *HttpErrorResponse {
	if (*r).Method != "POST" {
		return &HttpErrorResponse{
			Code:    InvalidRequestCode,
			Message: fmt.Sprintf("Invalid HTTP method. Expected POST, got %s", (*r).Method),
		}
	}

	if r.URL.Path != "/" {
		return &HttpErrorResponse{
			Code:    InvalidRequestCode,
			Message: fmt.Sprintf("Invalid HTTP path. Expected /, got %s", (*r).URL.Path),
		}
	}

	return nil
}

func (s *Server) Start() {
	var srv *http.Server

	srv = &http.Server{Addr: fmt.Sprintf(fmt.Sprintf("%s:%s", s.Host, s.Port))}

	http.HandleFunc("/", s.Handle)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), nil)

	if err != nil {
		fmt.Printf("An error occurred when starting Rose: %s\n", err.Error())
		fmt.Println("Shutting down the server gracefully...")

		// starting graceful shutdown of the server. all connections will be completed
		// it the context timeframe.
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			panic(err)
		}

		return
	}

	s.Server = srv
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	var err *HttpErrorResponse

	err = s.Validate(r)

	if err != nil {
		s.HandleError(w, r, err)

		return
	}
}

func (s *Server) HandleError(w http.ResponseWriter, r *http.Request, res *HttpErrorResponse) {
	SendResponse(res, w, 400)
}

func SendResponse(r *HttpErrorResponse, w http.ResponseWriter, status int) {
	var body []byte
	var jsonErr error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	body, jsonErr = json.Marshal(r)

	if jsonErr != nil {
		r = &HttpErrorResponse{
			Code:    SystemErrorCode,
			Message: "Could not create JSON from body. This is an internal unexpected error and it has been logged. Please, file an issue in https://github.com/MarioLegenda/rose/issues",
		}

		body, _ = json.Marshal(r)
	}

	_, err := w.Write(body)

	if err != nil {
		r = &HttpErrorResponse{
			Code:    SystemErrorCode,
			Message: "Could not write body to a response. This is an internal unexpected error and it has been logged. Please, file an issue in https://github.com/MarioLegenda/rose/issues",
		}

		body, _ = json.Marshal(r)

		_, _ = w.Write(body)
	}
}
