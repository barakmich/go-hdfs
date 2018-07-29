TRAVIS_TAG ?= $(shell git rev-parse HEAD)
ARCH = $(shell go env GOOS)-$(shell go env GOARCH)

all: hdfs

hdfs: clean deps
	go build -ldflags "-X main.version=$(TRAVIS_TAG)" ./cmd/hdfs

test: hdfs
	go test -v -race $(shell go list ./... | grep -v vendor)
	bats ./cmd/hdfs/test/*.bats

clean:
	rm -f ./hdfs

deps:
	go get -u github.com/colinmarc/hdfs
	go get -u github.com/spf13/cobra

.PHONY: clean test
