package main

import (
	"fmt"
	"os"
	"testing"
)


func TestSingleInsert(t *testing.T) {
	var s []byte
	var a *AppController
	var m *Metadata

	var appErr IError
	var runErr IError
	var appResult *AppResult

	s = []byte("sdčkfjalsčkjfdlsčakdfjlčk")

	m = &Metadata{
		Method: InsertMethodType,
		Data: &s,
		Id: "id",
	}

	a = &AppController{}
	appErr = a.Init()

	if appErr != nil {
		t.Errorf("TestSingleInsert: AppController failed to Init with message: %s", appErr.Error())

		return
	}

	runErr, appResult = a.Run(m)

	if runErr != nil {
		t.Errorf("TestSingleInsert: AppController::Run returned an error: %s", runErr.Error())

		return
	}

	if appResult.Status != "ok" {
		t.Errorf("TestSingleInsert: AppController::Run returned a non ok status but it should return ok")

		return
	}

	if appResult.Id != 0 {
		t.Errorf("TestSingleInsert: AppController::Run invalid Id returned on inisert. Got %d, expected %d", appResult.Id, 0)

		return
	}
}

func TestMultipleInsert(t *testing.T) {
	var s []byte
	var a *AppController
	var m *Metadata


	var appErr IError
	var appResult *AppResult
	var currId uint

	s = []byte("sdčkfjalsčkjfdlsčakdfjlčk")

	a = &AppController{}

	appErr = a.Init()

	if appErr != nil {
		t.Errorf("TestMultipleInsert: AppController failed to Init with message: %s", appErr.Error())

		return
	}

	for i := 0; i < 100000; i++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", i),
		}

		appErr, appResult = a.Run(m)

		if appErr != nil {
			t.Errorf("TestMultipleInsert: AppController::Run() returned an error: %s", appErr.Error())

			return
		}

		if appResult.Id != currId {
			t.Errorf("TestMultipleInsert: AppController::Run() there has been a discrepancy between generated id and counted id. Got %d, expected %d", appResult.Id, currId)

			return
		}

		currId++
	}
}

func TestDatabaseDirCreated(t *testing.T) {
	var m *Metadata
	var a *AppController
	var appErr IError

	a = &AppController{}

	appErr = a.Init()

	if appErr != nil {
		t.Errorf("TestDatabaseDirCreated: AppController failed to Init with message: %s", appErr.Error())

		return
	}

	m = &Metadata{
		Method: "insert",
		Data:   &[]byte{},
		Id: "validid",
	}

	err, _ := a.Run(m)

	if err != nil {
		t.Errorf("TestDatabaseDirCreated: ApplicationController::Run() returned an error: %s", err.Error())

		return
	}

	h := UserHomeDir()
	roseDb := fmt.Sprintf("%s/.rose_db", h)

	if _, err := os.Stat(roseDb); os.IsNotExist(err) {
		t.Errorf("TestDatabaseDirCreated: Database directory .rose_db was not created in %s", h)

		return
	}

	rmErr := os.RemoveAll(roseDb)
	if rmErr != nil {
		t.Errorf("TestDatabaseDirCreated: Database directory failed to remove")
	}
}

func TestInvalidMethod(t *testing.T) {
	var iv []string
	var m *Metadata
	var a *AppController

	var appErr IError

	a = &AppController{}
	appErr = a.Init()

	if appErr != nil {
		t.Errorf("TestInvalidMethod: failed to Init with message: %s", appErr.Error())

		return
	}


	iv = []string{"invalid1", "invalid2"}

	for i := 0; i < len(iv); i++ {
		m = &Metadata{
			Method: iv[i],
			Data:   &[]byte{},
			Id: "validid",
		}

		err, _ := a.Run(m)

		if err == nil {
			t.Errorf("TestInvalidMethod: ApplicationController::Run() should have returned an IError, nil returned")

			return
		}

		if err.Type() != HttpErrorType {
			t.Errorf("TestInvalidMethod: Invalid error type given. Expected %s, got %s", HttpErrorType, err.Type())
		}

		if err.GetCode() != HttpErrorCode {
			t.Errorf("TestInvalidMethod: Invalid error code given. Expected %d, got %d", HttpErrorCode, err.GetCode())
		}
	}
}

func TestInvalidId(t *testing.T) {
	var iv []string
	var m *Metadata
	var a *AppController

	var appErr IError

	a = &AppController{}

	appErr = a.Init()

	if appErr != nil {
		t.Errorf("TestInvalidId: AppController failed to Init with message: %s", appErr.Error())

		return
	}

	iv = []string{"insert", "read", "delete"}

	for i := 0; i < len(iv); i++ {
		m = &Metadata{
			Method: iv[i],
			Data:   &[]byte{},
			Id: "",
		}

		err, _ := a.Run(m)

		if err.GetCode() != HttpErrorCode {
			t.Errorf("TestInvalidId: Invalid error code given. Expected %d, got %d", HttpErrorCode, err.GetCode())
		}
	}
}

func TestValidMethod(t *testing.T) {
	var iv []string
	var m *Metadata
	var a *AppController

	var appErr IError

	a = &AppController{}

	appErr = a.Init()

	if appErr != nil {
		t.Errorf("TestValidMethod: AppController::Run() failed to Init with message: %s", appErr.Error())

		return
	}

	iv = []string{"insert", "read", "delete"}

	for i := 0; i < len(iv); i++ {
		m = &Metadata{
			Method: iv[i],
			Data:   &[]byte{},
			Id: "validid",
		}

		err, _ := a.Run(m)

		if err != nil {
			t.Errorf("TestValidMethod: ApplicationController::Run() returned an error: %s", err.Error())

			return
		}
	}
}
