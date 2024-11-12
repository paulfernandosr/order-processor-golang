package model

import (
	"net/http"
)

type Error struct {
	Status  int
	Message string
}

func NewBadRequestError(message string) *Error {
	return &Error{http.StatusBadRequest, message}
}

func NewNotFoundError(message string) *Error {
	return &Error{http.StatusNotFound, message}
}

func (err *Error) Error() string {
	return err.Message
}
