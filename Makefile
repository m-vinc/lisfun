
$(shell mkdir -p build/bin)

GO := go
GOPATH := $(shell go env GOPATH)
GOOS := $(shell go env GOOS)
GOARCH:= $(shell go env GOARCH)
VERSION := $(shell echo "dev")
OS_ARCHES := darwin_arm64

GORUN := go run -modfile=tools/go.mod
GOLANGCI_LINT := $(GORUN) github.com/golangci/golangci-lint/cmd/golangci-lint
GO_ARCH_LINT := $(GORUN) github.com/fe3dback/go-arch-lint

AIR := $(GORUN) github.com/air-verse/air
AIR_CONFIG := .air.toml

LISFUN_BINARY      := $(addprefix build/bin/lisfun_$(VERSION)_,$(OS_ARCHES))

GOFILES := $(shell find . -type f -name '*.go')
ASSETSFILES := $(shell find . -type f -name '*.css' -o -name '*.templ')

all: build

lint:
	$(GOLANGCI_LINT) run
	$(GO_ARCH_LINT) check
.PHONY: lint

build: lint $(GOFILES) $(ASSETSFILES)
	$(GO) build -o $(LISFUN_BINARY) ./cmd/lisfun/...

watch:
	$(AIR) -c $(AIR_CONFIG)
.PHONY: watch

clean:
	rm -rf build/
.PHONY: clean
