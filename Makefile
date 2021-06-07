.PHONY: test prepare

## Run all tests
test:
	go test -v -race -cover ./...

## Benchmark all tests
bench:
	go test -benchmem -bench=. ./...

## Prepare the codebase for commit by running `go fmt`, `golint` and `go mod tiny`
prepare: ${GO_FILES}
	go mod tidy
	go fmt ./...
	golint ./...
	golangci-lint run
