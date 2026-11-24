MAX_LINE_LENGTH := 100

OCI ?= podman-compose
RUN = $(OCI) run --rm -T checks
RUN_IT= $(OCI) run --rm checks
STAMP := .github/.checks-image

# Builds the checks image only when the Dockerfile or compose file changes
$(STAMP): Dockerfile compose.yml
	@$(OCI) build checks
	@touch $(STAMP)

.PHONY: build
build: $(STAMP)

.PHONY: test
test: build
	@clear || true
	@$(RUN) go test $(RACE) -count=1 -timeout 30s ./...
	@echo "✅ All tests passed!"

.PHONY: fmt
fmt: build
	@$(RUN) go fmt ./...
	@echo "✅ Code formatted!"

.PHONY: line
line: build
	@$(RUN) golines -w -m $(MAX_LINE_LENGTH) .
	@echo "✅ Lines limited!"

.PHONY: lint
lint: build
	@$(RUN) golangci-lint run
	@echo "✅ Linting passed!"

.PHONY: todo
todo: build
	@if $(RUN) grep -I --exclude="Makefile" --exclude-dir=".git" -r TODO .; then \
		echo "❌ Found TODOs" >&2; \
		exit 1; \
	fi

.PHONY: check
check: RACE=-race
check: test line fmt lint todo

.DEFAULT_GOAL := test
