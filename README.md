# gRPC Errors

[![BuildStatus Widget]][BuildStatus Result]
[![CodeCov Widget]][CodeCov Result]
[![GoReport Widget]][GoReport Status]
[![GoDoc Widget]][GoDoc]

[BuildStatus Result]: https://travis-ci.org/philip-bui/grpc-errors
[BuildStatus Widget]: https://travis-ci.org/philip-bui/grpc-errors.svg?branch=master

[CodeCov Result]: https://codecov.io/gh/philip-bui/grpc-errors
[CodeCov Widget]: https://codecov.io/gh/philip-bui/grpc-errors/branch/master/graph/badge.svg

[GoReport Status]: https://goreportcard.com/report/github.com/philip-bui/grpc-errors
[GoReport Widget]: https://goreportcard.com/badge/github.com/philip-bui/grpc-errors

[GoDoc]: https://godoc.org/github.com/philip-bui/grpc-errors
[GoDoc Widget]: https://godoc.org/github.com/philip-bui/grpc-errors?status.svg

Implementation of gRPC errors that return matching [status codes](https://github.com/grpc/grpc/blob/master/doc/statuscodes.md).

## Usage

```go
import (
	"github.com/philip-bui/grpc-errors"
)

func (s *Server) UnaryRequestExample(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return nil, errors.ErrInternalServer
}
```

## License

gRPC Errors is available under the MIT license. [See LICENSE](https://github.com/philip-bui/grpc-errors/blob/master/LICENSE) for details.
