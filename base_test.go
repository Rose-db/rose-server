package roseServer

import (
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"io/ioutil"
	"net"
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
		ginkgo.Fail(fmt.Sprintf("Cannot close connection: %s", err.Error()))
	}
}

func testReadUnixResponse(conn net.Conn) []uint8 {
	b, err := ioutil.ReadAll(conn)
	if err != nil {
		ginkgo.Fail(fmt.Sprintf("Cannot read server response: %s", err.Error()))
	}

	return b
}

func testCloseUnixConn(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		ginkgo.Fail(fmt.Sprintf("Cannot close connection: %s", err.Error()))
	}
}

func testWriteUnixServer(conn net.Conn, s []uint8) {
	if _, err := conn.Write(s); err != nil {
		ginkgo.Fail(fmt.Sprintf("Cannot write to unix server with message: %s", err.Error()))
	}
}

func testUnixConnect() net.Conn {
	conn, err := net.Dial("unix", "/tmp/rose.sock")

	if err != nil {
		ginkgo.Fail(fmt.Sprintf("Cannot connect to unix socket: %s", err.Error()))
	}

	return conn
}

func testCreateUDSWrite(m []uint8) []uint8 {
	s := socketRequest{
		Method:   "write",
		Metadata: m,
	}

	j, err := json.Marshal(s)

	if err != nil {
		ginkgo.Fail(fmt.Sprintf("Unable to marshal socketRequest: %s", err.Error()))
	}

	j = append(j, 10)

	return j
}

func testCreateSocketRequest(method string, data []uint8) []uint8 {
	s := socketRequest{
		Method:   methodType(method),
		Metadata: data,
	}

	j, err := json.Marshal(s)

	if err != nil {
		ginkgo.Fail(fmt.Sprintf("Unable to marshal socketRequest: %s", err.Error()))
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

func testAsJsonInterface(j interface{}) []uint8 {
	js, err := json.Marshal(j)

	if err != nil {
		ginkgo.Fail(fmt.Sprintf("Cannot marshal json with message: %s", err.Error()))
	}

	return js
}


