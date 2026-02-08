.PHONY: test
test:
	@clear || true
	go test ./...
	@echo "✅ All tests passed!"

.PHONY: cover
cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: fmt
fmt:
	@go fmt ./...
	@echo "✅ Code formatted!"

.PHONY: lint
lint:
	@golangci-lint run
	@echo "✅ Linting passed!"

.PHONY: todo
todo:
	@if grep -I --exclude="Makefile" --exclude-dir=".git" -r TODO .; then \
		echo "❌ Found TODOs" >&2; \
		exit 1; \
	fi

.PHONY: check
check: test fmt lint todo

.DEFAULT_GOAL := test
