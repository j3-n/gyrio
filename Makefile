GO ?= go
GOFMT ?= gofmt "-s"
PKL ?= "pkl-gen-go"
GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list ./...)

.phony: all
all:
	@$(GO) run cmd/cli/main.go

.phony: build
ifeq ($(v),1)
BUILD_FLAGS = -v
endif
build: clean
	@$(GO) build $(BUILD_FLAGS) -o build/program/gyrio cmd/cli/main.go

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

.phony: pkl-gen
pkl-gen:
	$(PKL) pkl/AppConfig.pkl
