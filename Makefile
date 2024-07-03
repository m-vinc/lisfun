
$(shell mkdir -p build/bin)

GO := go
GOPATH := $(shell go env GOPATH)
GOOS := $(shell go env GOOS)
GOARCH:= $(shell go env GOARCH)
VERSION := $(shell echo "dev")
OS_ARCHES := darwin_arm64

GORUN := go run -modfile=tools/go.mod
GOLANGCI_LINT := $(GORUN) github.com/golangci/golangci-lint/cmd/golangci-lint
ENT := $(GO) run ./gen/schema/generate/main.go
GO_ARCH_LINT := $(GORUN) github.com/fe3dback/go-arch-lint
GO_TEMPL := $(GORUN) github.com/a-h/templ/cmd/templ

ATLAS := atlas

AIR := $(GORUN) github.com/air-verse/air
AIR_CONFIG := .air.toml

LISFUN_BINARY      := $(addprefix build/bin/lisfun_$(VERSION)_,$(OS_ARCHES))

GOFILES := $(shell find . -type f -name '*.go')
ASSETSFILES := $(shell find . -type f -name '*.css' -o -name '*.templ')

all: build

gen: $(GOFILES)
	$(GO_TEMPL) generate
	$(ENT) ./gen/schema ./internal/db lisfun/internal/db
	npx tailwindcss -i ./internal/app/views/styles/main.css -o ./internal/app/assets/main.css
.PHONY: gen

lint:
	$(GOLANGCI_LINT) run
	$(GO_ARCH_LINT) check
.PHONY: lint

build: $(GOFILES) $(ASSETSFILES) gen lint
	$(GO) build -o $(LISFUN_BINARY) ./cmd/lisfun/...

watch:
	$(AIR) -c $(AIR_CONFIG)
.PHONY: watch

atlas_diff:
	$(ATLAS) migrate diff $(MIGRATION_NAME) \
  --dir "file://internal/db/migrate/migrations" \
  --to "ent://gen/schema" \
  --dev-url "docker://postgres/16/lisfun_atlas_migration?search_path=public"

atlas_apply:
	$(ATLAS) migrate apply \
  --dir "file://internal/db/migrate/migrations" \
  --url "${LISFUN_DATABASE_URL}"


clean:
	rm -rf build/
.PHONY: clean
