package main

import "fmt"

type Metadata struct {
	Method 	string
	Id 		string
	Data 	*[]byte
}

func (m *Metadata) Validate() IError {
	var v []string = []string{"insert", "read", "delete"}

	if !UtilHasString(m.Method, v) {
		return &HttpError{
			Code:    HttpErrorCode,
			Message: fmt.Sprintf("Method %s does not exist", m.Method),
		}
	}

	if m.Id == "" {
		return &HttpError{
			Code:    HttpErrorCode,
			Message: fmt.Sprintf("Id cannot be an empty string"),
		}
	}

	return nil
}


