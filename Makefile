GOPATH=$(shell go env GOPATH)

.PHONY: all
all:
	go run main.go

.PHONY: bin-deps
bin-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint	
lint:
	$(GOPATH)/bin/golangci-lint run	
