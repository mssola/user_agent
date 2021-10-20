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

EPOCH_COMMIT ?= 2046da2a89bd
.PHONY: git-validation
git-validation:
	@git-validation -q -range $(EPOCH_COMMIT)..HEAD -travis-pr-only=false

.PHONY: cilint
cilint:
	@golangci-lint run
