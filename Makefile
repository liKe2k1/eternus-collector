GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

default: build

workdir:
	mkdir -p workdir

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/eternus-collector .

test: test-all

test-all:
	@go test -v $(GOPACKAGES)