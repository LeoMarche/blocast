# Basic Makefile for Golang project
# Includes GRPC Gateway, Protocol Buffers
SERVICE		?= $(shell basename `go list`)
VERSION		?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(PWD)/.version 2> /dev/null || echo v0)
PACKAGE		?= $(shell go list)
PACKAGES	?= $(shell go list ./...)
FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Binaries
PROTOC		?= protoc

.PHONY: help clean fmt lint vet test test-cover generate-grpc build build-docker all

default: help

help: ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

install: ## build and install go application executable
	@go install -v ./...

env: ## Print useful environment variables to stdout
	@echo 'Current dir: ${CURDIR}'
	@echo 'Current service: ${SERVICE}'
	@echo 'Current package: ${PACKAGE}'
	@echo 'Current version: ${VERSION}'

build: ## builds for linux platform
	@export CGO_ENABLED=0 && export GOARCH=amd64 && export GOOS=linux && go build main.go

build-windows: ## builds for windows x64 platform
	@export CGO_ENABLED=0 && export GOARCH=amd64 && export GOOS=windows && go build main.go

clean: ## go clean
	@go clean

clean-all:  ## remove all generated artifacts and clean all build artifacts
	@go clean -i ./...

test-verbose:  ## tests with verbose activated
	@go test -v ./...

test: ## run tests
	@go test ./...

test-bench: ## run benchmark tests
	@go test -bench ./...

test-cover: ## Run test coverage and generate html report
	@rm -fr coverage
	@mkdir coverage
	@go list -f '{{if gt (len .TestGoFiles) 0}}"go test -covermode count -coverprofile {{.Name}}.coverprofile -coverpkg ./... {{.ImportPath}}"{{end}}' ./... | xargs -I {} bash -c {}
	@echo "mode: count" > coverage/cover.out
	@grep -h -v "^mode:" *.coverprofile >> "coverage/cover.out"
	@rm *.coverprofile
	@go tool cover -html=coverage/cover.out -o=coverage/cover.html

test-all: test test-bench test-cover clean-all ## run all tests

all: clean-all test-all build ## cleans, tests and build