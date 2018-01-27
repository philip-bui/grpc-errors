package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrResourceExhausted = &ResourceExhaustedError{
	"The resource is currently unavailable. Please try again in a short while",
}

type ResourceExhaustedError struct {
	string
}

func NewResourceExhaustedError(s string) *ResourceExhaustedError {
	return &ResourceExhaustedError{s}
}

func (e *ResourceExhaustedError) Error() string {
	return e.string
}

func (e *ResourceExhaustedError) GRPCStatus() *status.Status {
	return status.New(codes.ResourceExhausted, e.Error())
}
