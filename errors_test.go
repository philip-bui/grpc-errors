package errors

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestErrors(t *testing.T) {
	suite.Run(t, new(ErrorsSuite))
}

type ErrorsSuite struct {
	suite.Suite
}

func (s *ErrorsSuite) AssertErrorImplementation(err error) {
	// We could use an interfaced error, but decline to since it's only used for testing.
	s.Implements((*error)(nil), err)
	s.Implements((*interface{ GRPCStatus() *status.Status })(nil), err)
}

func (s *ErrorsSuite) TestAlreadyExistsError() {
	s.AssertErrorImplementation(ErrAlreadyExists)
	err := NewAlreadyExistsError(ErrAlreadyExists.string)
	s.Equal(ErrAlreadyExists.Error(), err.Error())
	s.Equal(ErrAlreadyExists.GRPCStatus(), err.GRPCStatus())
	s.Equal(err.GRPCStatus().Code(), codes.AlreadyExists)
}

func (s *ErrorsSuite) TestInternalServerError() {
	s.AssertErrorImplementation(ErrInternalServer)
	err := NewInternalServerError(ErrInternalServer.string)
	s.Equal(ErrInternalServer.Error(), err.Error())
	s.Equal(ErrInternalServer.GRPCStatus(), err.GRPCStatus())
	s.Equal(err.GRPCStatus().Code(), codes.Internal)
}

func (s *ErrorsSuite) TestInvalidArgumentError() {
	s.AssertErrorImplementation(ErrInvalidArgument)
	err := NewInvalidArgumentError(ErrInvalidArgument.string)
	s.Equal(ErrInvalidArgument.Error(), err.Error())
	s.Equal(ErrInvalidArgument.GRPCStatus(), err.GRPCStatus())
	s.Equal(err.GRPCStatus().Code(), codes.InvalidArgument)
}

func (s *ErrorsSuite) TestNotFoundError() {
	s.AssertErrorImplementation(ErrNotFound)
	err := NewNotFoundError(ErrNotFound.string)
	s.Equal(ErrNotFound.Error(), err.Error())
	s.Equal(ErrNotFound.GRPCStatus(), err.GRPCStatus())
	s.Equal(err.GRPCStatus().Code(), codes.NotFound)
}

func (s *ErrorsSuite) TestPermissionDeniedError() {
	s.AssertErrorImplementation(ErrPermissionDenied)
	err := NewPermissionDeniedError(ErrPermissionDenied.string)
	s.Equal(ErrPermissionDenied.Error(), err.Error())
	s.Equal(ErrPermissionDenied.GRPCStatus(), err.GRPCStatus())
	s.Equal(err.GRPCStatus().Code(), codes.PermissionDenied)
}

func (s *ErrorsSuite) TestResourceExhaustedError() {
	s.AssertErrorImplementation(ErrResourceExhausted)
	err := NewResourceExhaustedError(ErrResourceExhausted.string)
	s.Equal(ErrResourceExhausted.Error(), err.Error())
	s.Equal(ErrResourceExhausted.GRPCStatus(), err.GRPCStatus())
	s.Equal(err.GRPCStatus().Code(), codes.ResourceExhausted)
}

func (s *ErrorsSuite) TestUnauthenticatedError() {
	s.AssertErrorImplementation(ErrUnauthenticated)
	err := NewUnauthenticatedError(ErrUnauthenticated.string)
	s.Equal(ErrUnauthenticated.Error(), err.Error())
	s.Equal(ErrUnauthenticated.GRPCStatus(), err.GRPCStatus())
	s.Equal(err.GRPCStatus().Code(), codes.Unauthenticated)
}
