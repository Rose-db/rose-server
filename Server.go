package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HttpErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Server struct {
	Host string
	Port string
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
	http.HandleFunc("/", s.Handle)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), nil)

	if err != nil {
		fmt.Printf("An error occurred when starting Rose: %s\n", err.Error())

		return
	}
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	var res *HttpErrorResponse

	res = s.Validate(r)

	if res != nil {
		s.HandleError(w, r, res)

		return
	}
}

func (s *Server) HandleError(w http.ResponseWriter, r *http.Request, res *HttpErrorResponse) {
	SendResponse(res, w, 400)
}

func SendResponse(r *HttpErrorResponse, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	body, jsonErr := json.Marshal(r)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	_, err := w.Write(body)

	if err != nil {
		log.Fatal(err)
	}
}
