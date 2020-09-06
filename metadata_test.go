package main

import (
	"fmt"
	"os"
	"testing"
)

func TestInvalidMethod(t *testing.T) {
	var iv []string
	var m *Metadata
	var a *AppController

	a = &AppController{}


	iv = []string{"invalid1", "invalid2"}

	for i := 0; i < len(iv); i++ {
		m = &Metadata{
			Method: iv[i],
			Data:   &[]byte{},
		}

		err := a.Run(m)

		if err == nil {
			t.Errorf("ApplicationController::Run() should have returned an IError, nil returned")

			return
		}

		if err.Type() != HttpErrorType {
			t.Errorf("Invalid error type given. Expected %s, got %s", HttpErrorType, err.Type())
		}

		if err.GetCode() != HttpErrorCode {
			t.Errorf("Invalid error code given. Expected %d, got %d", HttpErrorCode, err.GetCode())
		}
	}
}

func TestValidMethod(t *testing.T) {
	var iv []string
	var m *Metadata
	var a *AppController

	a = &AppController{}

	iv = []string{"insert", "read", "delete"}

	for i := 0; i < len(iv); i++ {
		m = &Metadata{
			Method: iv[i],
			Data:   &[]byte{},
		}

		err := a.Run(m)

		if err != nil {
			t.Errorf("ApplicationController::Run() returned an error: %s", err.Error())

			return
		}
	}
}

func TestDatabaseDirCreated(t *testing.T) {
	var m *Metadata
	var a *AppController

	a = &AppController{}

	m = &Metadata{
		Method: "insert",
		Data:   &[]byte{},
	}

	err := a.Run(m)

	if err != nil {
		t.Errorf("ApplicationController::Run() returned an error: %s", err.Error())

		return
	}

	h := UserHomeDir()
	roseDb := fmt.Sprintf("%s/.rose_db", h)

	if _, err := os.Stat(roseDb); os.IsNotExist(err) {
		t.Errorf("Database directory .rose_db was not created in %s", h)

		return
	}

	rmErr := os.RemoveAll(roseDb)
	if rmErr != nil {
		t.Errorf("Database directory failed to remove")
	}
}