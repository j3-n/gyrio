GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list ./...)

.phony: all
all:
	@$(GO) run cmd/cli/main.go

.phony: build
build:
	@$(GO) build -o build/program/app cmd/cli/main.go

.phony: clean
clean:
	@rm -rf build
	@$(GO) clean

.phony: fmt
fmt:
	@$(GOFMT) -w $(GOFILES)

.phony: test
test:
	@$(GO) clean -testcache
	@$(GO) mod tidy
	@$(GO) test -cover $(GOMODULES)

.phony: update
update:
	@$(GO) get -u ./...
	@$(GO) mod tidy

.phony: info
info:
	@$(GO) vet $(GOMODULES)
	@$(GO) list $(GOMODULES)
	@$(GO) version