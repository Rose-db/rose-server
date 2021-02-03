package roseServer

import (
	"encoding/json"
	"fmt"
	"net"
	"rose/rose"
)

func createCollection(conn net.Conn, r *rose.Rose, req socketRequest) {
	roseErr := r.NewCollection(req.Metadata.(string))

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
	mtd := req.Metadata.(map[string]interface{})
	var m rose.WriteMetadata = rose.WriteMetadata{
		CollectionName: mtd["collectionName"].(string),
		Data:           mtd["data"],
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
	mtd := req.Metadata.(map[string]interface{})
	var m rose.ReadMetadata = rose.ReadMetadata{
		CollectionName: mtd["collectionName"].(string),
	}

	tmp := mtd["id"].(float64)
	id := int(tmp)

	m.ID = id

	rp := make(map[string]interface{})
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
	mtd := req.Metadata.(map[string]interface{})
	var m rose.DeleteMetadata = rose.DeleteMetadata{
		CollectionName: mtd["collectionName"].(string),
	}

	tmp := mtd["id"].(float64)
	id := int(tmp)

	m.ID = id

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

func replaceDocument(conn net.Conn, r *rose.Rose, req socketRequest) {
	var m rose.ReplaceMetadata

	err := json.Unmarshal([]uint8(req.Metadata.(string)), &m)

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

	res, roseErr := r.Replace(m)

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