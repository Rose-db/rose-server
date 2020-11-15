package roseServer

import (
	"fmt"
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
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		if ok := s.validate(r, "/read"); !ok {
			s.send404(w)

			return
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if ok := s.validate(r, "/delete"); !ok {
			s.send404(w)

			return
		}
	})

	http.HandleFunc("/size", func(w http.ResponseWriter, r *http.Request) {
		if ok := s.validate(r, "/size"); !ok {
			s.send404(w)

			return
		}
	})

	fmt.Printf("Starting Rose HTTP server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (s *httpServer) write(id string, data []uint8) {

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

func (s *httpServer) validate(r *http.Request, route string) bool {
	return r.URL.Path == route && r.Method == "POST"
}