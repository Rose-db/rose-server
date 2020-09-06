package main

import "fmt"

const SystemErrorType = "system_error"
const HttpErrorType = "http_error"

const HttpErrorCode = 1
const SystemErrorCode = 2

const InsertMethodType = "insert"
const DeleteMethodType = "delete"
const ReadMethodType = "read"

type IError interface {
	Error() string
	Type() string
	GetCode() int
	JSON() map[string]interface{}
}

type SystemError struct {
	Code int
	Message string
}

type HttpError struct {
	Code int
	Message string
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (e *SystemError) Type() string {
	return SystemErrorType
}

func (e *SystemError) GetCode() int {
	return SystemErrorCode
}

func (e *SystemError) JSON() map[string]interface{} {
	return map[string]interface{}{}
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (e *HttpError) Type() string {
	return HttpErrorType
}

func (e *HttpError) GetCode() int {
	return HttpErrorCode
}

func (e *HttpError) JSON() map[string]interface{} {
	return map[string]interface{}{}
}


