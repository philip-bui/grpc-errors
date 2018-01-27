package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUnauthenticated = &UnauthenticatedError{
	"You must be logged in to perform this action",
}

type UnauthenticatedError struct {
	string
}

func NewUnauthenticatedError(s string) *UnauthenticatedError {
	return &UnauthenticatedError{s}
}

func (e *UnauthenticatedError) Error() string {
	return e.string
}

func (e *UnauthenticatedError) GRPCStatus() *status.Status {
	return status.New(codes.Unauthenticated, e.Error())
}
