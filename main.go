package main

import (
	"fmt"
)

func main() {
	// Error message is handled in Boot()
	if ok := Boot(); !ok {
		return
	}

	server := &Server{
		Host: "127.0.0.1",
		Port: "8090",
	}

	fmt.Println("Rose server started. Listening to incoming requests")

	server.Start()
}

func Boot() bool {
	var app *AppController
	var err error
	var errStream chan IError

	fmt.Println("Booting Rose cache server...")

	app = &AppController{}
	errStream = app.Init(true)

	err = <- errStream

	if err != nil {
		fmt.Printf("An error occurred when starting Rose: %s\nExiting", err.Error())

		return false
	}

	return true
}