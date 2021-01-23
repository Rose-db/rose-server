package roseServer

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	Error() string
	Type() string
	GetCode() int
	JSON() []uint8
}

type unixSocketError struct {
	Code int
	Message string
}

func (e *unixSocketError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (e *unixSocketError) Type() string {
	return UnixSocketErrorType
}

func (e *unixSocketError) GetCode() int {
	return UnixSocketErrorCode
}

func (e *unixSocketError) JSON() []uint8 {
	return errToJson(e)
}



type requestError struct {
	Code int
	Message string
}

func (e *requestError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (e *requestError) Type() string {
	return RequestErrorType
}

func (e *requestError) GetCode() int {
	return RequestErrorCode
}

func (e *requestError) JSON() []uint8 {
	return errToJson(e)
}




type systemError struct {
	Code int
	Message string
}

func (e *systemError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (e *systemError) Type() string {
	return SystemErrorType
}

func (e *systemError) GetCode() int {
	return SystemErrorCode
}

func (e *systemError) JSON() []uint8 {
	return errToJson(e)
}

func errToJson(e Error) []uint8 {
	j := map[string]interface{}{
		"type": e.Type(),
		"code": e.GetCode(),
		"message": e.Error(),
	}

	b, _ := json.Marshal(j)

	return b
}


