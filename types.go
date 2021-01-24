package roseServer

import "rose/rose"

type socketRequest struct {
	Method methodType `json:"method"`
	Metadata []uint8  `json:"metadata"`
}

type socketResponse struct {
	Method methodType `json:"method"`
	Status int `json:"status"`
	Error interface{} `json:"error"`
	Data *rose.AppResult `json:"data"`
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
type ErrorType string

const SystemErrorType ErrorType = "system_error"
const RequestErrorType ErrorType = "request_error"

// application error codes
type ErrorCode int

const InvalidRequestDataErrorCode ErrorCode = 1
const InvalidRequestMethodErrorCode ErrorCode = 2
const InvalidMetadataErrorCode ErrorCode = 3
const InvalidStartUpErrorCode ErrorCode = 4

const OperationSuccessCode = 1
const OperationFailedCode = 0

type ServerType string

// server types
const UnixSocketServer ServerType = "uds"
