package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrPermissionDenied = &PermissionDeniedError{
	"You do not have permission to perform this action",
}

type PermissionDeniedError struct {
	string
}

func NewPermissionDeniedError(s string) *PermissionDeniedError {
	return &PermissionDeniedError{s}
}

func (e *PermissionDeniedError) Error() string {
	return e.string
}

func (e *PermissionDeniedError) GRPCStatus() *status.Status {
	return status.New(codes.PermissionDenied, e.Error())
}
