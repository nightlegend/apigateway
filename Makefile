GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")

all: install

install: deps
	govendor sync

