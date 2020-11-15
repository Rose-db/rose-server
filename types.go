package roseServer

type readRequest struct {
	Id 		string `json:"id"`
}

type writeRequest struct {
	Id 		string `json:"id"`
	Data 	string `json:"data"`
}

type deleteRequest struct {
	Id 		string `json:"id"`
	Data 	string `json:"data"`
}

type response struct {
	Method string	`json:"method"`
	Status string	`json:"status"`
	Reason string	`json:"reason"`
	Result string	`json:"result"`
}

type errorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Server interface {
	Start()
}

// error types
const systemErrorType = "system_error"

// application error codes
const systemErrorCode = 2

// server types
const httpServerType = "http"
