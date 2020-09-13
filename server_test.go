package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func testSendRequest(method string, id string, data string, b *testing.B) []byte {
	var body []byte
	var err error
	var response *http.Response

	m := make(map[string]string)
	m["method"] = method
	m["id"] = id
	m["data"] = data

	body, _ = json.Marshal(m)

	request, err := http.NewRequest("POST", "http://127.0.0.1:8090/", bytes.NewBuffer(body))

	if err != nil {
		b.Errorf("%s: An error occurred creating request: %s", testGetBenchmarkName(b), err.Error())

		return nil
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accepts", "application/json")

	client := &http.Client{
		Timeout: time.Second * time.Duration(100),
	}

	response, err = client.Do(request)

	if err != nil {
		b.Errorf("%s: An error occurred making request: %s", testGetBenchmarkName(b), err.Error())

		return nil
	}

	if response.StatusCode != 200 {
		b.Errorf("%s: Invalid status returned. Expected 200, got %d", testGetBenchmarkName(b), response.StatusCode)

		return nil
	}

	body, err = ioutil.ReadAll(response.Body)

	if err != nil {
		b.Errorf("%s: An error occurred making request: %s", testGetBenchmarkName(b), err.Error())

		return nil
	}

	defer response.Body.Close()

	if len(body) == 0 {
		b.Errorf("%s: Empty body returned in sending request. It should have returned something", testGetBenchmarkName(b))

		return nil
	}

	return body
}

func testCreateTestServer(method string, url string, body io.Reader, withApp bool, t *testing.T) *httptest.ResponseRecorder {
	var rr *httptest.ResponseRecorder
	var server *Server
	var rErr error
	var req *http.Request
	var handler http.HandlerFunc
	var app *AppController

	req, rErr = http.NewRequest(method,  url, body)
	if rErr != nil {
		t.Errorf("%s: Request could not be created with message %s", testGetTestName(t), rErr.Error())

		return nil
	}

	if withApp {
		app = &AppController{}
		app.Init(false)
	}

	server = &Server{
		Host: "127.0.0.1",
		Port: "8090",
		App: app,
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
		rr = testCreateTestServer(methods[i], "/", nil, false, t)

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

	rr = testCreateTestServer("POST", "/invalid", nil, false, t)

	if status := rr.Code; status != 400 {
		t.Errorf("%s: Invalid status returned. Expected %d, got %d", testGetTestName(t), 400, status)
	}

	expected = fmt.Sprintf(`{"code":3,"message":"Invalid HTTP path. Expected /, got /invalid"}`)
	given = rr.Body.String()

	if expected != given {
		t.Errorf("%s: Invalid response body returned. Expected %s, got %s", testGetTestName(t), expected, given)
	}
}

func TestBasicValidHttpInsert(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var expected, given string

	m := make(map[string]string)
	m["method"] = InsertMethodType
	m["id"] = "bogus id"
	m["data"] = "bogus data"

	body, _ := json.Marshal(m)

	rr = testCreateTestServer("POST", "/", bytes.NewReader(body), true, t)

	if status := rr.Code; status != 200 {
		t.Errorf("%s: Invalid status returned. Expected %d, got %d; Response: %s", testGetTestName(t), 200, status, rr.Body.String())

		return
	}

	expected = `{"id":0,"method":"insert","status":"ok","reason":"","result":""}`
	given = rr.Body.String()

	if expected != given {
		t.Errorf("%s: Invalid result returned. Expected %s, got %s", testGetTestName(t), expected, given)

		return
	}
}

