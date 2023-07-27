package model

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

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

func WrapGormError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Error(http.StatusNotFound, err.Error(), "db entry not found")
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return Error(http.StatusConflict, err.Error(), "unique constraint violation")
	}
	return Error(http.StatusInternalServerError, err.Error(), "internal server error")
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
