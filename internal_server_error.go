package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInternalServer = &InternalServerError{
	"There was an internal server error. Please try again",
}

type InternalServerError struct {
	string
}

func NewInternalServerError(s string) *InternalServerError {
	return &InternalServerError{s}
}

func (e *InternalServerError) Error() string {
	return e.string
}

func (e *InternalServerError) GRPCStatus() *status.Status {
	return status.New(codes.Internal, e.Error())
}
