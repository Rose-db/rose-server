package roseServer

import (
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"io/ioutil"
	"net"
	"rose/rose"
	"testing"
)

var GomegaRegisterFailHandler = gomega.RegisterFailHandler
var GinkgoFail = ginkgo.Fail
var GinkgoRunSpecs = ginkgo.RunSpecs
var GinkgoBeforeSuite = ginkgo.BeforeSuite
var GinkgoAfterSuite = ginkgo.AfterSuite
var GinkgoDescribe = ginkgo.Describe
var GinkgoIt = ginkgo.It

func TestRose(t *testing.T) {
	GomegaRegisterFailHandler(GinkgoFail)
	GinkgoRunSpecs(t, "Rose Server Suite")
}

func testCloseUnixWriteConn(conn net.Conn) {
	err := conn.(*net.UnixConn).CloseWrite()
	if err != nil {
		panic(err)
	}
}

func testReadUnixResponse(conn net.Conn) []uint8 {
	b, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	return b
}

func testCloseUnixConn(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}

func testWriteUnixServer(conn net.Conn, s []uint8) {
	if _, err := conn.Write(s); err != nil {
		panic(err)
	}
}

func testUnixConnect() net.Conn {
	conn, err := net.Dial("unix", "/tmp/rose.sock")

	if err != nil {
		panic(err)
	}

	return conn
}

func testCreateCollection(collName string) {
	conn := testUnixConnect()

	req := testCreateSocketRequest("createCollection", []uint8(collName))

	testWriteUnixServer(conn, req)

	testCloseUnixWriteConn(conn)
	b := testReadUnixResponse(conn)

	var res socketResponse
	err := json.Unmarshal(b, &res)

	if err != nil {
		panic(err)
	}

	testCloseUnixConn(conn)
}

func testWrite(conn net.Conn, m rose.WriteMetadata) {
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	req := testCreateSocketRequest("write", b)

	testWriteUnixServer(conn, req)

	testCloseUnixWriteConn(conn)

	b = testReadUnixResponse(conn)

	var res socketResponse
	err = json.Unmarshal(b, &res)

	if err != nil {
		panic(err)
	}

	gomega.Expect(res.Data.Status).To(gomega.Equal(rose.OkResultStatus))
}

func testRead(conn net.Conn, m rose.ReadMetadata, writeData string) socketResponse {
	b, err := json.Marshal(m)

	if err != nil {
		gomega.Expect(err).To(gomega.BeNil())
	}

	req := testCreateSocketRequest("read", b)

	testWriteUnixServer(conn, req)

	testCloseUnixWriteConn(conn)

	b = testReadUnixResponse(conn)

	var res socketResponse
	err = json.Unmarshal(b, &res)

	if err != nil {
		panic(err)
	}

	gomega.Expect(res.Data.Status).To(gomega.Equal(rose.FoundResultStatus))

	return res
}

func testCreateSocketRequest(method string, data []uint8) []uint8 {
	s := socketRequest{
		Method:   methodType(method),
		Metadata: string(data),
	}

	j, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	j = append(j, 10)

	return j
}

func testAsJson(j string) []uint8 {
	js, err := json.Marshal(j)

	if err != nil {
		ginkgo.Fail(fmt.Sprintf("Cannot marshal json with message: %s", err.Error()))
	}

	return js
}