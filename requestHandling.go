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
		if ok := writeUDSError(
			conn,
			fmt.Sprintf("Unable to read request body: %s", err.Error()),
			"",
			InvalidRequestDataErrorCode,
			RequestErrorType);
		!ok {
			return
		}

		return
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
			return
		}

		return
	}

	if !methodTypesImpl.IncludesType(req.Method) {
		if ok := writeUDSError(
			conn,
			fmt.Sprintf("Invalid method '%s'. Expected one of '%s'", req.Method, methodTypesImpl.String()),
			"",
			InvalidRequestMethodErrorCode,
			RequestErrorType);
		!ok {
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
			Data: nil,
		}); !ok {
			// write to log

			return
		}

		return
	} else if req.Method == write {
		var wm rose.WriteMetadata

		err := json.Unmarshal(req.Metadata, &wm)

		if err != nil {
			if ok := writeUDSError(
				conn,
				fmt.Sprintf("Cannot read WRITE request metadata with message: %s", err.Error()),
				string(write),
				InvalidMetadataErrorCode,
				RequestErrorType);
			!ok {
				return
			}

			return
		}

		res, roseErr := r.Write(wm)

		if roseErr != nil {
			if ok := writeRoseError(conn, roseErr); !ok {
				return
			}

			return
		}

		if ok := writeSuccessResponse(conn, socketResponse{
			Method: req.Method,
			Status: OperationSuccessCode,
			Data: res,
		}); !ok {
			// write to log

			return
		}

		return
	} else if req.Method == read {
		var rm rose.ReadMetadata

		err := json.Unmarshal(req.Metadata, &rm)

		if err != nil {
			if ok := writeUDSError(
				conn,
				fmt.Sprintf("Cannot read READ request metadata with message: %s", err.Error()),
				string(read),
				InvalidMetadataErrorCode,
				RequestErrorType);
			!ok {
				return
			}

			return
		}

		rp := make(map[string]interface{})

		rm.Data = &rp

		_, roseErr := r.Read(rm)

		if roseErr != nil {
			if ok := writeRoseError(conn, roseErr); !ok {
				return
			}

			return
		}
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



