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
	var m rose.WriteMetadata

	err := json.Unmarshal(req.Metadata, &m)

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

	res, roseErr := r.Write(m)

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
	var m rose.ReadMetadata

	err := json.Unmarshal(req.Metadata, &m)

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
	m.Data = &rp

	res, roseErr := r.Read(m)

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
		ReadData: m.Data,
	}); !ok {
		// write to log

		return
	}
}

func deleteDocument(conn net.Conn, r *rose.Rose, req socketRequest) {
	var m rose.DeleteMetadata

	err := json.Unmarshal(req.Metadata, &m)

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

	res, roseErr := r.Delete(m)

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