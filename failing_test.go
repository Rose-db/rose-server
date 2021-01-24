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

		gomega.Expect(errRes["Status"].(float64)).To(gomega.Equal(float64(OperationFailedCode)))
		gomega.Expect(errRes["Method"]).To(gomega.Equal(""))

		data := errRes["Data"].(map[string]interface{})

		gomega.Expect(data["Type"]).To(gomega.Equal(RequestErrorType))
		gomega.Expect(data["Code"].(float64)).To(gomega.Equal(float64(RequestErrorCode)))
		gomega.Expect(data["Message"]).To(gomega.Equal("Code: 3, Message: Invalid method invalid. Expected one of createCollection, write, read, delete, replace, query"))

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

		gomega.Expect(errRes["Status"].(float64)).To(gomega.Equal(float64(OperationFailedCode)))
		gomega.Expect(errRes["Method"]).To(gomega.Equal(""))

		data := errRes["Data"].(map[string]interface{})

		gomega.Expect(data["Type"]).To(gomega.Equal(SystemErrorType))
		gomega.Expect(data["Code"].(float64)).To(gomega.Equal(float64(SystemErrorCode)))
		gomega.Expect(data["Message"]).To(gomega.Equal("Code: 2, Message: Unable to read request body: EOF"))

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

		gomega.Expect(errRes["Status"].(float64)).To(gomega.Equal(float64(OperationFailedCode)))
		gomega.Expect(errRes["Method"]).To(gomega.Equal(""))

		data := errRes["Data"].(map[string]interface{})

		gomega.Expect(data["Type"]).To(gomega.Equal(SystemErrorType))
		gomega.Expect(data["Code"].(float64)).To(gomega.Equal(float64(SystemErrorCode)))
		gomega.Expect(data["Message"]).To(gomega.Equal("Code: 2, Message: Cannot unpack request body: unexpected end of JSON input"))

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
		gomega.Expect(errRes["Status"].(float64)).To(gomega.Equal(float64(OperationFailedCode)))
		gomega.Expect(errRes["Method"]).To(gomega.Equal(""))

		data := errRes["Data"].(map[string]interface{})

		gomega.Expect(data["Type"]).To(gomega.Equal(SystemErrorType))
		gomega.Expect(data["Code"].(float64)).To(gomega.Equal(float64(SystemErrorCode)))
		gomega.Expect(data["Message"]).To(gomega.Equal("Code: 2, Message: Cannot unpack request body: json: cannot unmarshal string into Go value of type roseServer.socketRequest"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because of invalid write request metadata", func() {
		conn := testUnixConnect()

		req := testCreateSocketRequest("write", []uint8("invalid values"))

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		errRes := make(map[string]interface{})

		err := json.Unmarshal(b, &errRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(errRes["Status"].(float64)).To(gomega.Equal(float64(OperationFailedCode)))
		gomega.Expect(errRes["Method"]).To(gomega.Equal(string(write)))

		data := errRes["Data"].(map[string]interface{})

		gomega.Expect(data["Type"]).To(gomega.Equal(RequestErrorType))
		gomega.Expect(data["Code"].(float64)).To(gomega.Equal(float64(RequestErrorCode)))
		gomega.Expect(data["Message"]).To(gomega.Equal("Code: 3, Message: Cannot read WRITE request metadata with message: invalid character 'i' looking for beginning of value"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because of invalid read request metadata", func() {
		conn := testUnixConnect()

		req := testCreateSocketRequest("read", []uint8("myColl"))

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		errRes := make(map[string]interface{})

		err := json.Unmarshal(b, &errRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(errRes["Status"].(float64)).To(gomega.Equal(float64(OperationFailedCode)))
		gomega.Expect(errRes["Method"]).To(gomega.Equal(string(read)))

		data := errRes["Data"].(map[string]interface{})

		gomega.Expect(data["Type"]).To(gomega.Equal(RequestErrorType))
		gomega.Expect(data["Code"].(float64)).To(gomega.Equal(float64(RequestErrorCode)))
		gomega.Expect(data["Message"]).To(gomega.Equal("Code: 3, Message: Cannot read READ request metadata with message: invalid character 'm' looking for beginning of value"))

		testCloseUnixConn(conn)
	})
})


