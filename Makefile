GO ?= go
GO_SRC = $(shell find . -name \*.go)

.DEFAULT: build
all: test lint

.PHONY: test
test:
	@$(GO) test

.PHONY: bench
bench:
	@$(GO) test -bench=.

.PHONY: lint
lint: git-validation cilint

EPOCH_COMMIT ?= 834b6d4d9e84
.PHONY: git-validation
git-validation:
	@git-validation -v -D -range $(EPOCH_COMMIT)..HEAD

.PHONY: cilint
cilint:
	@golangci-lint run
