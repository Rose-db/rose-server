package roseServer

import (
	"fmt"
	"net"
	"rose/rose"
)

func newUDSServer() Server {
	return &UDSServer{}
}

type UDSServer struct {}

func (u *UDSServer) Start() Error {
	l, err := net.Listen("unix", "/tmp/rose.sock")

	if err != nil {
		return &unixSocketError{
			Code:    UnixSocketErrorCode,
			Message: fmt.Sprintf("Failed listening to unix domain socket with message: %s", err.Error()),
		}
	}

	r, err := rose.New(true)

	if err != nil {
		return &systemError{
			Code:    SystemErrorCode,
			Message: fmt.Sprintf("Rose failed to create: %s", err.Error()),
		}
	}

	fmt.Println("\u001B[32mSocket server is ready to accept requests!\033[0m")
	for {
		conn, err := l.Accept()
		if err != nil {
			return &unixSocketError{
				Code:    UnixSocketErrorCode,
				Message: fmt.Sprintf("Failed accepting a request on unix domain socket with message: %s", err.Error()),
			}
		}

		go runRequest(conn, r)
	}
}
