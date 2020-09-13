package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testCreateTestServer(method string, url string, body io.Reader, t *testing.T) *httptest.ResponseRecorder {
	var rr *httptest.ResponseRecorder
	var server *Server
	var rErr error
	var req *http.Request
	var handler http.HandlerFunc

	req, rErr = http.NewRequest(method,  url, body)
	if rErr != nil {
		t.Errorf("%s: Request could not be created with message %s", testGetTestName(t), rErr.Error())

		return nil
	}

	server = &Server{
		Host: "127.0.0.1",
		Port: "8090",
		App: nil,
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(server.Handle)

	handler.ServeHTTP(rr, req)

	return rr
}

func TestInvalidHttpMethod(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var expected string
	var given string
	var methods []string

	methods = []string{
		"GET",
		"PUT",
		"DELETE",
		"OPTIONS",
		"TRACE",
		"HEAD",
		"CONNECT",
		"PATCH",
	}

	for i := range methods {
		rr = testCreateTestServer(methods[i], "/", nil, t)

		if status := rr.Code; status != 400 {
			t.Errorf("%s: Invalid status returned. Expected %d, got %d", testGetTestName(t), 400, status)
		}

		expected = fmt.Sprintf(`{"code":3,"message":"Invalid HTTP method. Expected POST, got %s"}`, methods[i])
		given = rr.Body.String()

		if expected != given {
			t.Errorf("%s: Invalid response body returned. Expected %s, got %s", testGetTestName(t), expected, given)
		}
	}
}

func TestInvalidHttpPath(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var expected, given string

	rr = testCreateTestServer("POST", "/invalid", nil, t)

	if status := rr.Code; status != 400 {
		t.Errorf("%s: Invalid status returned. Expected %d, got %d", testGetTestName(t), 400, status)
	}

	expected = fmt.Sprintf(`{"code":3,"message":"Invalid HTTP path. Expected /, got /invalid"}`)
	given = rr.Body.String()

	if expected != given {
		t.Errorf("%s: Invalid response body returned. Expected %s, got %s", testGetTestName(t), expected, given)
	}
}

