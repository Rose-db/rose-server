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

		var sockRes socketResponse
		err := json.Unmarshal(b, &sockRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(sockRes.Status).To(gomega.Equal(OperationFailedCode))
		gomega.Expect(string(sockRes.Method)).To(gomega.Equal(""))

		gomega.Expect(sockRes.Data).To(gomega.BeNil())

		sockErr := sockRes.Error.(map[string]interface{})
		
		gomega.Expect(sockErr["Type"].(string)).To(gomega.Equal(string(RequestErrorType)))
		gomega.Expect(sockErr["Code"].(float64)).To(gomega.Equal(float64(InvalidRequestMethodErrorCode)))
		gomega.Expect(sockErr["Message"]).To(gomega.Equal("Invalid method 'invalid'. Expected one of 'createCollection, write, read, delete, replace, query'"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because the server cannot read an empty request body", func() {
		conn := testUnixConnect()

		testWriteUnixServer(conn, []uint8{})

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		var sockRes socketResponse
		err := json.Unmarshal(b, &sockRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(sockRes.Status).To(gomega.Equal(OperationFailedCode))
		gomega.Expect(string(sockRes.Method)).To(gomega.Equal(""))

		gomega.Expect(sockRes.Data).To(gomega.BeNil())

		sockErr := sockRes.Error.(map[string]interface{})

		gomega.Expect(sockErr["Type"].(string)).To(gomega.Equal(string(RequestErrorType)))
		gomega.Expect(sockErr["Code"]).To(gomega.Equal(float64(InvalidRequestDataErrorCode)))
		gomega.Expect(sockErr["Message"]).To(gomega.Equal("Unable to read request body: EOF"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because the server cannot read request body", func() {
		conn := testUnixConnect()

		testWriteUnixServer(conn, []uint8{'\n'})

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		var sockRes socketResponse
		err := json.Unmarshal(b, &sockRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(sockRes.Status).To(gomega.Equal(OperationFailedCode))
		gomega.Expect(string(sockRes.Method)).To(gomega.Equal(""))

		gomega.Expect(sockRes.Data).To(gomega.BeNil())

		sockErr := sockRes.Error.(map[string]interface{})

		gomega.Expect(sockErr["Type"].(string)).To(gomega.Equal(string(RequestErrorType)))
		gomega.Expect(sockErr["Code"].(float64)).To(gomega.Equal(float64(InvalidRequestDataErrorCode)))
		gomega.Expect(sockErr["Message"]).To(gomega.Equal("Cannot unpack request body: unexpected end of JSON input"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because the server cannot read invalid request body", func() {
		conn := testUnixConnect()

		s := testAsJson("something")
		s = append(s, '\n')

		testWriteUnixServer(conn, s)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		var sockRes socketResponse
		err := json.Unmarshal(b, &sockRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(sockRes.Status).To(gomega.Equal(OperationFailedCode))
		gomega.Expect(string(sockRes.Method)).To(gomega.Equal(""))

		gomega.Expect(sockRes.Data).To(gomega.BeNil())

		sockErr := sockRes.Error.(map[string]interface{})

		gomega.Expect(sockErr["Type"].(string)).To(gomega.Equal(string(RequestErrorType)))
		gomega.Expect(sockErr["Code"].(float64)).To(gomega.Equal(float64(InvalidRequestDataErrorCode)))
		gomega.Expect(sockErr["Message"]).To(gomega.Equal("Cannot unpack request body: json: cannot unmarshal string into Go value of type roseServer.socketRequest"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because of invalid write request metadata", func() {
		conn := testUnixConnect()

		req := testCreateSocketRequest("write", []uint8("invalid values"))

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		var sockRes socketResponse
		err := json.Unmarshal(b, &sockRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(sockRes.Status).To(gomega.Equal(OperationFailedCode))
		gomega.Expect(string(sockRes.Method)).To(gomega.Equal(string(write)))

		sockErr := sockRes.Error.(map[string]interface{})

		gomega.Expect(sockErr["Type"].(string)).To(gomega.Equal(string(RequestErrorType)))
		gomega.Expect(sockErr["Code"].(float64)).To(gomega.Equal(float64(InvalidMetadataErrorCode)))
		gomega.Expect(sockErr["Message"]).To(gomega.Equal("Cannot read WRITE request metadata with message: invalid character 'i' looking for beginning of value"))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should fail the request because of invalid read request metadata", func() {
		conn := testUnixConnect()

		req := testCreateSocketRequest("read", []uint8("myColl"))

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		var sockRes socketResponse
		err := json.Unmarshal(b, &sockRes)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal error response: %s", err.Error()))
		}

		gomega.Expect(sockRes.Status).To(gomega.Equal(OperationFailedCode))
		gomega.Expect(string(sockRes.Method)).To(gomega.Equal(string(read)))

		sockErr := sockRes.Error.(map[string]interface{})

		gomega.Expect(sockErr["Type"].(string)).To(gomega.Equal(string(RequestErrorType)))
		gomega.Expect(sockErr["Code"].(float64)).To(gomega.Equal(float64(InvalidMetadataErrorCode)))
		gomega.Expect(sockErr["Message"]).To(gomega.Equal("Cannot read READ request metadata with message: invalid character 'm' looking for beginning of value"))

		testCloseUnixConn(conn)
	})
})


