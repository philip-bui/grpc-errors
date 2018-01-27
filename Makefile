PORT?=8000
PACKAGE:=github.com/philip-bui/grpc-zerolog
COVERAGE:=coverage.txt
godoc:
	echo "localhost:${PORT}/pkg/${PACKAGE}"
	godoc -http=:${PORT}

.PHONY: test

test:
	go test -race -coverprofile=${COVERAGE} -covermode=atomic

coverage:
	go tool cover -html=${COVERAGE}
