package errors

import "net/http"

// apiError is an error wrapper that also
// holds information about response status codes.
type apiError struct {
	error
	statusCode int
}

// HTTPErrorStatusCode returns a status code.
func (e apiError) HTTPErrorStatusCode() int {
	return e.statusCode
}

// NewErrorWithStatusCode allows you to associate
// a specific HTTP Status Code to an error.
// The Server will take that code and set
// it as the response status.
func NewErrorWithStatusCode(err error, code int) error {
	return apiError{err, code}
}

// NewBadRequestError creates a new API error
// that has the 400 HTTP status code associated to it.
func NewBadRequestError(err error) error {
	return NewErrorWithStatusCode(err, http.StatusBadRequest)
}

// NewRequestNotFoundError creates a new API error
// that has the 404 HTTP status code associated to it.
func NewRequestNotFoundError(err error) error {
	return NewErrorWithStatusCode(err, http.StatusNotFound)
}

// NewRequestConflictError creates a new API error
// that has the 409 HTTP status code associated to it.
func NewRequestConflictError(err error) error {
	return NewErrorWithStatusCode(err, http.StatusConflict)
}
