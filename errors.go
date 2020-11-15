package roseServer

import "fmt"

type Error interface {
	Error() string
	Type() string
	Code() int
	JSON() map[string]interface{}
}

type systemError struct {
	Message string
}

func (e *systemError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code(), e.Message)
}

func (e *systemError) Type() string {
	return systemErrorType
}

func (e *systemError) Code() int {
	return SystemErrorCode
}

func (e *systemError) JSON() map[string]interface{} {
	return map[string]interface{}{}
}



