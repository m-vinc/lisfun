
$(shell mkdir -p build/bin)

GO := go
GOPATH := $(shell go env GOPATH)
GOOS := $(shell go env GOOS)
GOARCH:= $(shell go env GOARCH)
VERSION := $(shell echo "dev")
OS_ARCHES := darwin_arm64

AIR := $(addprefix $(GOPATH)/,bin/air)
AIR_CONFIG := .air.toml

LISFUN_BINARY      := $(addprefix build/bin/lisfun_$(VERSION)_,$(OS_ARCHES))

GOFILES := $(shell find . -type f -name '*.go')
ASSETSFILES := $(shell find . -type f -name '*.css' -o -name '*.templ')

deps:
	go install github.com/air-verse/air@latest

lint:
	golangci-lint run

build: lint $(GOFILES) $(ASSETSFILES)
	$(GO) build -o $(LISFUN_BINARY) ./cmd/lisfun/...

watch:
	$(AIR) -c $(AIR_CONFIG)
.PHONY: watch

clean:
	rm -rf build/
.PHONY: clean
