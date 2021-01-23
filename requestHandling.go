package roseServer

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"rose/rose"
)

func runRequest(conn net.Conn, r *rose.Rose) {
	defer conn.Close()

	body := &bytes.Buffer{}
	s, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		if ok := writeUDSSystemError(conn, fmt.Sprintf("Unable to read request body: %s", err.Error())); !ok {
			return
		}

		return
	}

	body.Write(s)

	var req socketRequest
	if err := json.Unmarshal(body.Bytes(), &req); err != nil {
		if ok := writeUDSSystemError(conn, fmt.Sprintf("Cannot unpack request body: %s", err.Error())); !ok {
			return
		}

		return
	}

	if !methodTypesImpl.IncludesType(req.Method) {
		if ok := writeUDSRequestError(conn, fmt.Sprintf("Invalid method %s. Expected one of %s", req.Method, methodTypesImpl.String())); !ok {
			return
		}

		return
	}

	if req.Method == createCollection {
		roseErr := r.NewCollection(string(req.Metadata))

		if roseErr != nil {
			if ok := writeRoseError(conn, roseErr); !ok {
				// write to log
				return
			}

			return
		}

		if ok := writeSuccessResponse(conn, socketResponse{
			Method: req.Method,
			Status: OperationSuccessCode,
			Result: nil,
		}); !ok {
			// write to log

			return
		}

		return
	}
}

func writeUDSSystemError(conn net.Conn, msg string) bool {
	_, err := conn.Write((&systemError{
		Code:    SystemErrorCode,
		Message: msg,
	}).JSON())

	return err == nil
}

func writeUDSRequestError(conn net.Conn, msg string) bool {
	_, err := conn.Write((&requestError{
		Code:    RequestErrorCode,
		Message: msg,
	}).JSON())

	return err == nil
}

func writeRoseError(conn net.Conn, err rose.Error) bool {
	_, e := conn.Write(err.JSON())

	return e == nil
}

func writeSuccessResponse(conn net.Conn, res socketResponse) bool {
	b, _ := json.Marshal(res)

	_, e := conn.Write(b)

	return e == nil
}


