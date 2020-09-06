package main

type AppController struct {}

func (a *AppController) Run(m *Metadata) IError {
	err := m.Validate()

	if err != nil {
		return err
	}

	fsErr := CreateDbIfNotExists()

	if fsErr != nil {
		return err
	}

	db := &Database{InternalDb: nil}

	db.Init()

	if m.Method == InsertMethodType {
		db.Insert("sfdjasdlfjk", m.Data)
	}

	// LoadDatabase()

	return nil
}
