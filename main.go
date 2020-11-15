package main

import (
	"fmt"
	"rose/rose"
)

func main() {
	r, err := rose.New(true)

	if err != nil {
		panic(err)
	}

	server := &Server{
		Host: "127.0.0.1",
		Port: "8090",
		App: r,
	}

	fmt.Println(fmt.Sprintf("Rose server started on %s:%s. Listening to incoming requests", server.Host, server.Port))

	server.Start()
}