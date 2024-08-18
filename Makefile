GOPATH ?= $(shell go env GOPATH)
GO_MOD_DIRS := $(shell find . -type f -name "go.mod" -exec dirname {} \;)
YELLOW := "\033[0;33m"
#RED := "\033[0;31m"
BLUE := "\033[0;34m"
RED := $(shell tput setaf 1)
GREEN := $(shell tput setaf 2)
YELLOW := $(shell tput setaf 3)
NC := $(shell tput sgr0)

#
# Format
#

.PHONY fmt:
fmt: py-fmt go-fmt

.PHONY py-fmt:
py-fmt:
	@echo "$(YELLOW)Formatting Python code with ruff...$(NC)"
	python -m ruff check --fix-only
	python -m ruff format

.PHONY go-fmt:
go-fmt:
	@echo "$(YELLOW)Formatting Go code...$(NC)"
	@for dir in $(GO_MOD_DIRS); do \
		echo Formatting $$dir...; \
		gofmt -s -w .; \
		cd $$dir && golangci-lint run --fix; \
		cd - >/dev/null; \
	done
	@echo Done.

#
# Lint
#

.PHONY lint:
lint: py-lint go-lint

.PHONY py-lint:
py-lint:
	@echo "$(YELLOW)Linting Python code with ruff...$(NC)"
	python -m ruff check --exit-non-zero-on-fix
	python -m ruff check --preview --exit-non-zero-on-fix
	python -m ruff format --check

.PHONY go-lint:
go-lint: ensure-gopath
	@echo "$(YELLOW)Linting Go code...$(NC)"
	@test -e $(GOPATH)/bin/golangci-lint || \
		echo Installing linters... && \
		(curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.54.2)

	@for dir in $(GO_MOD_DIRS); do \
		echo Linting $$dir...; \
		cd $$dir && golangci-lint run -v; \
		cd - >/dev/null; \
	done
	@echo Done.

#
# Build
#

.PHONY build-go-fasthttp:
build-go-fasthttp:
	cd services/go_fasthttp && make build

.PHONY build-py-flask:
build-py-flask:
	cd services/py_flask && make build

#
# Environment
#

.PHONY: ensure-gopath
ensure-gopath:
ifndef GOPATH
	$(error GOPATH must be set)
endif
