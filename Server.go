package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type HttpRequestModel struct {
	Method 	string `json:"method"`
	Id 		string `json:"id"`
	Data 	string `json:"data"`
}

type HttpResponseModel struct {
	Id uint 		`json:"id"`
	Method string	`json:"method"`
	Status string	`json:"status"`
	Reason string	`json:"reason"`
	Result string	`json:"result"`
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
	var appErr IError
	var metadata *Metadata
	var appResult *AppResult
	var responseModel *HttpResponseModel

	metadata = s.CreateMetadata(w, r)

	if metadata == nil {
		return
	}

	appErr, appResult = s.App.Run(metadata)

	if appErr != nil {
		var httpErr *HttpErrorResponse

		httpErr = &HttpErrorResponse{
			Code:    appErr.GetCode(),
			Message: appErr.Error(),
		}

		s.SendError(httpErr, w)

		return
	}

	responseModel = &HttpResponseModel{
		Id:     appResult.Id,
		Method: appResult.Method,
		Status: appResult.Status,
		Reason: "",
		Result: appResult.Result,
	}

	s.SendSuccess(responseModel, w)
}

func (s *Server) SendError(res *HttpErrorResponse, w http.ResponseWriter) {
	var jsonErr error
	var body []byte

	body, jsonErr = json.Marshal(res)

	if jsonErr != nil {
		var r *HttpErrorResponse

		r = &HttpErrorResponse{
			Code:    SystemErrorCode,
			Message: "There has been an error but could not create the body for it. This is an internal unexpected error and it has been logged. Please, file an issue in https://github.com/MarioLegenda/rose/issues",
		}

		body, _ = json.Marshal(r)
	}

	s.SendResponse(body, 400, w)
}

func (s *Server) SendSuccess(res *HttpResponseModel, w http.ResponseWriter) {
	var jsonErr error
	var body []byte

	body, jsonErr = json.Marshal(res)

	if jsonErr != nil {
		s.SendError(&HttpErrorResponse{
			Code:    SystemErrorCode,
			Message: "There has been an error but could not create the body for it. This is an internal unexpected error and it has been logged. Please, file an issue in https://github.com/MarioLegenda/rose/issues",
		}, w)

		return
	}

	s.SendResponse(body, 200, w)
}

func (s *Server) ReadBody(b io.Reader, method string, w http.ResponseWriter) (*HttpErrorResponse, []byte) {
	var r *HttpErrorResponse

	body, err := ioutil.ReadAll(b)

	if err != nil {
		r = &HttpErrorResponse{
			Code:    SystemErrorCode,
			Message: "Could not create JSON from body. This is an internal unexpected error and it has been logged. Please, file an issue in https://github.com/MarioLegenda/rose/issues",
		}

		return r, nil
	}

	return nil, body
}

func (s *Server) SendResponse(body []byte, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err := w.Write(body)

	if err != nil {
		var r *HttpErrorResponse

		r = &HttpErrorResponse{
			Code:    SystemErrorCode,
			Message: "Could not write body to a response. This is an internal unexpected error and it has been logged. Please, file an issue in https://github.com/MarioLegenda/rose/issues",
		}

		body, _ = json.Marshal(r)

		_, _ = w.Write(body)
	}
}

func (s *Server) CreateMetadata(w http.ResponseWriter, r *http.Request) *Metadata {
	var httpErr *HttpErrorResponse
	var jErr error
	var model HttpRequestModel
	var data []byte

	httpErr = s.Validate(r)

	if httpErr != nil {
		s.SendError(httpErr, w)

		return nil
	}

	httpErr, body := s.ReadBody(r.Body, r.Method, w)

	if httpErr != nil {
		s.SendError(httpErr, w)

		return nil
	}

	jErr = json.Unmarshal(body, &model)

	if jErr != nil {
		httpErr = &HttpErrorResponse{
			Code:    SystemErrorCode,
			Message: fmt.Sprintf("Could not read body. Error : %s", jErr.Error()),
		}

		s.SendError(httpErr, w)

		return nil
	}

	data = []byte(model.Data)

	return &Metadata{
		Method: model.Method,
		Id:     model.Id,
		Data:   &data,
	}
}
