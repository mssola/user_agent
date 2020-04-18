GO ?= go
GO_SRC = $(shell find . -name \*.go)

.DEFAULT: build
build: $(GO_SRC)
	@$(GO) build

#
# Unit & integration tests.
#

.PHONY: test
test:
	@$(GO) test

.PHONY: bench
bench:
	@$(GO) test -bench=.

#
# Validation tools.
#

EPOCH_COMMIT ?= 2046da2a89bd
.PHONY: git-validation
git-validation:
ifeq (, $(shell which git-validation 2> /dev/null))
	@echo "You don't have 'git-validation' installed, consider installing it (see the CONTRIBUTING.md file)."
else
	@git-validation -q -range $(EPOCH_COMMIT)..HEAD -travis-pr-only=false
endif

.PHONY: validate-go
validate-go:
	@which gofmt >/dev/null 2>/dev/null || (echo "ERROR: gofmt not found." && false)
	@test -z "$$(gofmt -s -l . | grep -vE '^vendor/' | tee /dev/stderr)"
	@chmod +x ci/lint.sh
	@./ci/lint.sh
# Go 1.4.x does not have go vet, so there is no point on trying.
ifneq ("$(TRAVIS_GO_VERSION)","1.4.x")
	@go doc cmd/vet >/dev/null 2>/dev/null || (echo "ERROR: go vet not found." && false)
	@test -z "$$(go vet . | grep -v vendor | tee /dev/stderr)"
endif

.PHONY: validate
validate: git-validation validate-go

#
# Travis-CI
#

.PHONY: ci
ci: build validate test
