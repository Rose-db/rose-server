package main

type ReadRequest struct {
	Id 		string `json:"id"`
}

type WriteRequest struct {
	Id 		string `json:"id"`
	Data 	string `json:"data"`
}

type DeleteRequest struct {
	Id 		string `json:"id"`
	Data 	string `json:"data"`
}

type Response struct {
	Method string	`json:"method"`
	Status string	`json:"status"`
	Reason string	`json:"reason"`
	Result string	`json:"result"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Server interface {
	Start()
}
