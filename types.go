package roseServer

type socketRequest struct {
	Method methodType `json:"method"`
	Metadata []uint8  `json:"metadata"`
}

type socketResponse struct {
	Method methodType `json:"method"`
	Status int `json:"status"`
	Result interface{} `json:"result"`
}

type Server interface {
	Start() Error
}

type methodTypes struct {
	types []methodType
}

var methodTypesImpl methodTypes = methodTypes{
	types: []methodType{
		createCollection,
		write,
		read,
		delete,
		replace,
		query,
	},
}

func (m methodTypes) IncludesType(a methodType) bool {
	for _, b := range m.types {
		if a == b {
			return true
		}
	}

	return false
}

func (m methodTypes) String() string {
	s := ""
	for i, b := range m.types {
		s += string(b)

		if i != len(m.types) - 1 {
			s += ", "
		}
	}

	return s
}

type methodType string

const createCollection methodType = "createCollection"
const write methodType = "write"
const read methodType = "read"
const delete methodType = "delete"
const replace methodType = "replace"
const query methodType = "query"

// error types
const UnixSocketErrorType = "unix_socket_error"
const SystemErrorType = "system_error"
const RequestErrorType = "request_error"

// application error codes
const UnixSocketErrorCode = 1
const SystemErrorCode = 2
const RequestErrorCode = 3

const OperationSuccessCode = 1

type ServerType string

// server types
const UnixSocketServer ServerType = "uds"
