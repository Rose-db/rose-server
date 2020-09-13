package main

import (
	"fmt"
)

func main() {
	var app *AppController
	// Error message is handled in Boot()
	if app = Boot(); app == nil {
		return
	}

	server := &Server{
		Host: "127.0.0.1",
		Port: "8090",
		App: app,
	}

	fmt.Println(fmt.Sprintf("Rose server started on %s:%s. Listening to incoming requests", server.Host, server.Port))

	server.Start()
}

func Boot() *AppController {
	var app *AppController
	var err error
	var errStream chan IError

	fmt.Println("Booting Rose cache server...")

	app = &AppController{}
	errStream = app.Init(true)

	err = <- errStream

	if err != nil {
		fmt.Printf("An error occurred when starting Rose: %s\nExiting", err.Error())

		return nil
	}

	return app
}