package roseServer

import (
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = GinkgoDescribe("Failing tests", func() {
	GinkgoIt("Should fail because of invalid request method", func() {
		conn := testUnixConnect()

		req := testCreateSocketRequest("invalid", []uint8("myColl"))

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		errRes := make(map[string]interface{})

		err := json.Unmarshal(b, &errRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(errRes["type"]).To(gomega.Equal(RequestErrorType))
		gomega.Expect(errRes["code"].(float64)).To(gomega.Equal(float64(RequestErrorCode)))
		gomega.Expect(errRes["message"]).To(gomega.Equal("Code: 3, Message: Invalid method invalid. Expected one of createCollection, write, read, delete, replace, query"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because the server cannot read and empty request body", func() {
		conn := testUnixConnect()

		testWriteUnixServer(conn, []uint8{})

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		errRes := make(map[string]interface{})

		err := json.Unmarshal(b, &errRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(errRes["type"]).To(gomega.Equal(SystemErrorType))
		gomega.Expect(errRes["code"].(float64)).To(gomega.Equal(float64(SystemErrorCode)))
		gomega.Expect(errRes["message"]).To(gomega.Equal("Code: 2, Message: Unable to read request body: EOF"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because the server cannot read request body", func() {
		conn := testUnixConnect()

		testWriteUnixServer(conn, []uint8{'\n'})

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		errRes := make(map[string]interface{})

		err := json.Unmarshal(b, &errRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(errRes["type"]).To(gomega.Equal(SystemErrorType))
		gomega.Expect(errRes["code"].(float64)).To(gomega.Equal(float64(SystemErrorCode)))
		gomega.Expect(errRes["message"]).To(gomega.Equal("Code: 2, Message: Cannot unpack request body: unexpected end of JSON input"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because the server cannot read invalid request body", func() {
		conn := testUnixConnect()

		s := testAsJson("something")
		s = append(s, '\n')

		testWriteUnixServer(conn, s)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		errRes := make(map[string]interface{})

		err := json.Unmarshal(b, &errRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(errRes["type"]).To(gomega.Equal(SystemErrorType))
		gomega.Expect(errRes["code"].(float64)).To(gomega.Equal(float64(SystemErrorCode)))
		gomega.Expect(errRes["message"]).To(gomega.Equal("Code: 2, Message: Cannot unpack request body: json: cannot unmarshal string into Go value of type roseServer.socketRequest"))

		testCloseUnixConn(conn)
	})
})


