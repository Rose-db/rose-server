package main

type AppController struct {
	Database *Database
}

type AppResult struct {
	Id uint
	Method string
	Status string
	Reason string
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
			Status: "ok",
			Reason: "",
		}
	} else if m.Method == DeleteMethodType {
		return nil, &AppResult{
			Id:     1,
			Method: m.Method,
			Status: "ok",
			Reason: "",
		}
	} else if m.Method == ReadMethodType {
		return nil, &AppResult{
			Id:     1,
			Method: m.Method,
			Status: "ok",
			Reason: "",
		}
	}

	panic("Internal rose error. Unreachable code reached. None of the methods have executed but one should have.")
}

func (a *AppController) Init() IError {
	var fsErr IError

	fsErr = CreateDbIfNotExists()

	if fsErr != nil {
		return fsErr
	}

	a.Database = &Database{}

	a.Database.Init()

	return nil
}
