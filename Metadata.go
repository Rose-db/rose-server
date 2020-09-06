package main

import "fmt"

type Metadata struct {
	Method string
	Data *[]byte
}

func (m *Metadata) Validate() IError {
	var v []string = []string{"insert", "read", "delete"}

	if !UtilHasString(m.Method, v) {
		return &HttpError{
			Code:    0,
			Message: fmt.Sprintf("Method %s does not exist", m.Method),
		}
	}

	return nil
}


