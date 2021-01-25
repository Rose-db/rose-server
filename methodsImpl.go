package roseServer

import (
	"encoding/json"
	"fmt"
	"net"
	"rose/rose"
)

func createCollection(conn net.Conn, r *rose.Rose, req socketRequest) {
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
}

func createDocument(conn net.Conn, r *rose.Rose, req socketRequest) {
	var wm rose.WriteMetadata

	err := json.Unmarshal(req.Metadata, &wm)

	if err != nil {
		if ok := writeUDSError(
			conn,
			fmt.Sprintf("Cannot read WRITE request metadata with message: %s", err.Error()),
			string(writeMethod),
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
}

func readDocument(conn net.Conn, r *rose.Rose, req socketRequest) {
	var rm rose.ReadMetadata

	err := json.Unmarshal(req.Metadata, &rm)

	if err != nil {
		if ok := writeUDSError(
			conn,
			fmt.Sprintf("Cannot read READ request metadata with message: %s", err.Error()),
			string(readMethod),
			InvalidMetadataErrorCode,
			RequestErrorType);
			!ok {
			return
		}

		return
	}

	rp := ""
	rm.Data = &rp

	res, roseErr := r.Read(rm)

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
		ReadData: rm.Data,
	}); !ok {
		// write to log

		return
	}
}