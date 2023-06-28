package exceptions

// BadRequestError represents a bad request error.
type BadRequestError struct {
	message string
}

func (e BadRequestError) Error() string {
	return e.message
}

// NotFoundError represents a not found error.
type NotFoundError struct {
	message string
}

func (e NotFoundError) Error() string {
	return e.message
}
