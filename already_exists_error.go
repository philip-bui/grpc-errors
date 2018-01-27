package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrAlreadyExists = &AlreadyExistsError{
	"The resource already exists",
}

type AlreadyExistsError struct {
	string
}

func NewAlreadyExistsError(s string) *AlreadyExistsError {
	return &AlreadyExistsError{s}
}

func (e *AlreadyExistsError) Error() string {
	return e.string
}

func (e *AlreadyExistsError) GRPCStatus() *status.Status {
	return status.New(codes.AlreadyExists, e.Error())
}
