package roseServer

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testCreateTestServer(method string, url string, body io.Reader, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method,  url, body)
	if err != nil {
		t.Errorf("Request could not be created with message %s", err.Error())

		return nil
	}

	server := &Server{
		Host: "127.0.0.1",
		Port: "8090",
		App: nil,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Handle)

	handler.ServeHTTP(rr, req)

	return rr
}


