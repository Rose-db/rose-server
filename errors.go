package roseServer

type Error interface {
	Error() string
	GetType() ErrorType
	GetCode() ErrorCode
	Map() map[string]interface{}
}




type serverError struct {
	Type ErrorType
	Code ErrorCode
	Message string
}

func (e *serverError) Error() string {
	return e.Message
}

func (e *serverError) GetType() ErrorType {
	return e.Type
}

func (e *serverError) GetCode() ErrorCode {
	return e.Code
}

func (e *serverError) Map() map[string]interface{} {
	return errToJson(e)
}




func errToJson(e Error) map[string]interface{} {
	return map[string]interface{}{
		"Type": e.GetType(),
		"Code": e.GetCode(),
		"Message": e.Error(),
	}
}


