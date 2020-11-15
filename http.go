package roseServer

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type httpServer struct {
	Host string
	Port string
	Sdk *sdk
}

func newHttpServer() *httpServer {
	sdk, err := newSdk()

	if err != nil {
		panic(err)
	}

	return &httpServer{
		Host: "",
		Port: "",
		Sdk:  sdk,
	}
}

func (s *httpServer) Start() {
	http.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		if ok := s.validate(r, "/write"); !ok {
			s.send404(w)

			return
		}

		body, err := s.readBody(r.Body)

		if err != nil {
			s.send400(w, err)

			return
		}
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		if ok := s.validate(r, "/read"); !ok {
			s.send404(w)

			return
		}

		body, err := s.readBody(r.Body)

		if err != nil {
			s.send400(w, err)

			return
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if ok := s.validate(r, "/delete"); !ok {
			s.send404(w)

			return
		}

		body, err := s.readBody(r.Body)

		if err != nil {
			s.send400(w, err)

			return
		}
	})

	http.HandleFunc("/size", func(w http.ResponseWriter, r *http.Request) {
		if ok := s.validate(r, "/size"); !ok {
			s.send404(w)

			return
		}

		body, err := s.readBody(r.Body)

		if err != nil {
			s.send400(w, err)

			return
		}
	})

	fmt.Printf("Starting Rose HTTP server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (s *httpServer) send404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	_, err := w.Write([]uint8{})

	// gracefully shutdown the server here since a response must be written
	if err != nil {
		log.Fatal(err)
	}
}

func (s *httpServer) send400(w http.ResponseWriter, e Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	r := &errorResponse{
		Code:    systemErrorCode,
		Message: e.Error(),
	}

	body, err := json.Marshal(r)

	if err != nil {
		_, err := w.Write([]uint8{})

		if err != nil {
			// gracefully shutdown the server here, MUST write
			log.Fatal(err)
		}
	}

	_, err = w.Write(body)

	if err != nil {
		_, err := w.Write([]uint8{})

		if err != nil {
			// gracefully shutdown the server here, MUST write
			log.Fatal(err)
		}
	}
}

func (s *httpServer) validate(r *http.Request, route string) bool {
	return r.URL.Path == route && r.Method == "POST"
}

func (s *httpServer) readBody(b io.Reader) ([]uint8, Error) {
	body, err := ioutil.ReadAll(b)

	if err != nil {
		return []uint8{}, &systemError{
			Message: "Cannot read request body",
		}
	}

	return body, nil
}