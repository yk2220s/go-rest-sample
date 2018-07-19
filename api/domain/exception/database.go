package exception

import "net/http"

// DatabaseError is error when error occur at database.
type DatabaseError struct {
	message    string
	statusCode int
}

// NewDatabaseError make DatabaseError
func NewDatabaseError(message string) DatabaseError {
	return DatabaseError{message, http.StatusInternalServerError}
}

func (e DatabaseError) Error() string {
	return e.message
}

// StatusCode return status code.
func (e DatabaseError) StatusCode() int {
	return e.statusCode
}
