package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInvalidArgument = &InvalidArgumentError{
	"There is an issue with your request",
}

type InvalidArgumentError struct {
	string
}

func NewInvalidArgumentError(s string) *InvalidArgumentError {
	return &InvalidArgumentError{s}
}

func (e *InvalidArgumentError) Error() string {
	return e.string
}

func (e *InvalidArgumentError) GRPCStatus() *status.Status {
	return status.New(codes.InvalidArgument, e.Error())
}
