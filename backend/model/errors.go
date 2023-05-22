package model

type HttpError struct {
	statusCode int
	details    string
	message    string
}

func Error(status int, details, message string) HttpError {
	return HttpError{
		statusCode: status,
		details:    details,
		message:    message,
	}
}

func (e HttpError) Error() string {
	return e.details
}

func (e HttpError) Status() int {
	return e.statusCode
}

func (e HttpError) Message() string {
	return e.message
}
