package error

import (
	"net/http"
)

type appError struct {
	cause      error  `json:"-"`
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	ErrorCode  string `json:"code"`
}

func (e appError) Error() string {
	if e.cause != nil {
		return e.Message + ": " + e.cause.Error()
	}
	return e.Message
}

func (e appError) Cause() error {
	return e.cause
}

func WithError(cause error, with appError) error {
	return appError{
		cause:      cause,
		StatusCode: with.StatusCode,
		Message:    with.Message,
		ErrorCode:  with.ErrorCode,
	}
}

func Cast(err error) appError {
	ae, ok := err.(appError)
	if ok {
		return ae
	}
	return InternalServerError
}

var InternalServerError = appError{
	StatusCode: http.StatusInternalServerError,
	Message:    "internal server error.",
	ErrorCode:  "INTERNAL_SERVER_ERROR",
}

var InvalidRequestParameters = appError{
	StatusCode: http.StatusBadRequest,
	Message:    "invalid request parameters.",
	ErrorCode:  "INVALID_REQUEST_PARAMETERS",
}

var DataNotFound = appError{
	StatusCode: http.StatusNotFound,
	Message:    "data not found.",
	ErrorCode:  "DATA_NOT_FOUND",
}

var NotFound = appError{
	StatusCode: http.StatusNotFound,
	Message:    "not found.",
	ErrorCode:  "NOT_FOUND",
}
