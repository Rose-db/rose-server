package roseServer

import (
	"fmt"
)

func NewServer(t ServerType) (Server, Error) {
	if t == UnixSocketServer {
		return newUDSServer(), nil
	}

	return nil, &systemError{
		Code:    SystemErrorCode,
		Message: fmt.Sprintf("Unknown server type %s", t),
	}
}
