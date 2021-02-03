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

	req, ok := validateRequest(conn)

	if !ok {
		return
	}

	if req.Method == createCollectionMethod {
		createCollection(conn, r, req)

		return
	} else if req.Method == writeMethod {
		createDocument(conn, r, req)

		return
	} else if req.Method == readMethod {
		readDocument(conn, r, req)

		return
	} else if req.Method == deleteMethod {
		deleteDocument(conn, r, req)
	} else if req.Method == replaceMethod {
		replaceDocument(conn, r, req)
	}
}

func writeUDSError(conn net.Conn, msg string, method string, code ErrorCode, t ErrorType) bool {
	reqErr := (&serverError{
		Type: t,
		Code:    code,
		Message: msg,
	}).Map()

	res := socketResponse{
		Method: methodType(method),
		Status: OperationFailedCode,
		Error:  reqErr,
		Data:   nil,
	}

	b, _ := json.Marshal(res)

	_, err := conn.Write(b)

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

func validateRequest(conn net.Conn) (socketRequest, bool) {
	body := &bytes.Buffer{}
	s, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		if ok := writeUDSError(
			conn,
			fmt.Sprintf("Unable to read request body: %s", err.Error()),
			"",
			InvalidRequestDataErrorCode,
			RequestErrorType);
			!ok {
			return socketRequest{}, false
		}

		return socketRequest{}, false
	}

	body.Write(s)

	var req socketRequest
	if err := json.Unmarshal(body.Bytes(), &req); err != nil {
		if ok := writeUDSError(
			conn,
			fmt.Sprintf("Cannot unpack request body: %s", err.Error()),
			"",
			InvalidRequestDataErrorCode,
			RequestErrorType);
			!ok {
			return socketRequest{}, false
		}

		return socketRequest{}, false
	}

	if !methodTypesImpl.IncludesType(req.Method) {
		if ok := writeUDSError(
			conn,
			fmt.Sprintf("Invalid method '%s'. Expected one of '%s'", req.Method, methodTypesImpl.String()),
			"",
			InvalidRequestMethodErrorCode,
			RequestErrorType);
			!ok {
			return socketRequest{}, false
		}

		return socketRequest{}, false
	}

	return req, true
}



