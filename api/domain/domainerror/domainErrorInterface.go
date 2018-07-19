package domainerror

// DomainError is error occured at domain objects.
type DomainError interface {
	error
	StatusCode() int
}
