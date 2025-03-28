.PHONY: test
test:
	go test -cover -timeout 10s ./...
	go vet ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: push
push:
	make test
	git pull origin HEAD
	git push origin HEAD

.DEFAULT_GOAL := test
