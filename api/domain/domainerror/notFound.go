package domainerror

import "net/http"

// NotFound is error when resource is not found
type NotFound struct {
	message    string
	statusCode int
}

// NewNotFound make NotFound
func NewNotFound() NotFound {
	return NotFound{"resource is not found.", http.StatusNotFound}
}

func (e NotFound) Error() string {
	return e.message
}

// StatusCode return status code.
func (e NotFound) StatusCode() int {
	return e.statusCode
}
