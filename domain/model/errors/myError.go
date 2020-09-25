package errors

type ErrorType string

type MyError struct {
	StatusCode int
	Message    string
	ErrorType  ErrorType
}

func (e MyError) Error() string {
	return e.Message
}
