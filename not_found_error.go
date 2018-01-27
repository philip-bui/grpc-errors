package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNotFound = &NotFoundError{
	"The resource could not be found",
}

type NotFoundError struct {
	string
}

func NewNotFoundError(s string) *NotFoundError {
	return &NotFoundError{s}
}

func (e *NotFoundError) Error() string {
	return e.string
}

func (e *NotFoundError) GRPCStatus() *status.Status {
	return status.New(codes.NotFound, e.Error())
}
