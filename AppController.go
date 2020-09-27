package main

import "fmt"

type AppController struct {
	Database *Database
}

type AppResult struct {
	Id uint
	Method string
	Status string
	Reason string
	Result string
}

func (a *AppController) Run(m *Metadata) (IError, *AppResult) {
	var vErr IError
	var idx uint

	vErr = m.Validate()

	if vErr != nil {
		return vErr, nil
	}

	if m.Method == InsertMethodType {
		idx = a.Database.Insert(m.Id, m.Data)

		return nil, &AppResult{
			Id:     idx,
			Method: m.Method,
			Status: FoundResultStatus,
		}
	} else if m.Method == DeleteMethodType {
		return nil, &AppResult{
			Id:     1,
			Method: m.Method,
			Status: FoundResultStatus,
		}
	} else if m.Method == ReadMethodType {
		var res *DbReadResult
		var err *DbReadError
		res, err = a.Database.Read(m.Id)

		if err != nil {
			return nil, &AppResult{
				Id:     0,
				Method: m.Method,
				Status: NotFoundResultStatus,
				Reason: err.Error(),
				Result: "",
			}
		}

		return nil, &AppResult{
			Id:     res.Idx,
			Method: m.Method,
			Status: FoundResultStatus,
			Result: res.Result,
		}
	}

	panic("Internal rose error. Unreachable code reached. None of the methods have executed but one should have.")
}

func (a *AppController) Init(log bool) chan IError {
	var fsStream chan string
	var errStream chan IError

	fsStream = make(chan string)
	errStream = make(chan IError)

	go CreateDbIfNotExists(fsStream, errStream)

	for msg := range fsStream {
		if log {
			fmt.Println(msg)
		}
	}

	a.Database = &Database{}

	a.Database.Init()

	return errStream
}
