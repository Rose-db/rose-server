package roseServer

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	Error() string
	Type() string
	GetCode() int
	JSON(method string) []uint8
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

func (e *unixSocketError) JSON(method string) []uint8 {
	return errToJson(e, method)
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

func (e *requestError) JSON(method string) []uint8 {
	return errToJson(e, method)
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

func (e *systemError) JSON(method string) []uint8 {
	return errToJson(e, method)
}

func errToJson(e Error, method string) []uint8 {
	j := map[string]interface{}{
		"Status": OperationFailedCode,
		"Method": method,
		"Data": map[string]interface{}{
			"Type": e.Type(),
			"Code": e.GetCode(),
			"Message": e.Error(),
		},
	}

	b, _ := json.Marshal(j)

	return b
}


