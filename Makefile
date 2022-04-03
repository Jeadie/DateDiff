default: build

LATEST_TAG := $(shell git tag | tail -n 1)
GO_OS := $(shell go env GOOS)
GO_ARCH := $(shell go env GOARCH)

build:
	go build .

test:
	test -z `go fmt ./...`
	go test ./...

builddockerbinary:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o datediff_docker .
	chmod +x datediff_docker

release: test build

CI: test

download:
	curl -OL https://github.com/Jeadie/DateDiff/releases/download/$(LATEST_TAG)/datediff-$(LATEST_TAG)-$(GO_OS)-$(GO_ARCH).tar.gz
	tar -xzf  datediff-$(LATEST_TAG)-$(GO_OS)-$(GO_ARCH).tar.gz
