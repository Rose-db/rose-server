package main

// error types
const SystemErrorType = "system_error"
const HttpErrorType = "http_error"
const DatabaseIntegrityViolationErrorType = "db_integrity_violation"

// application error codes
const HttpErrorCode = 1
const SystemErrorCode = 2
const InvalidRequestCode = 3
const DbIntegrityViolationCode = 4

// database error codes
const InvalidReadErrorCode = 5

// result status
const FoundResultStatus = "ok"
const NotFoundResultStatus = "not_found"

// method types
const InsertMethodType = "insert"
const DeleteMethodType = "delete"
const ReadMethodType = "read"

