test:
	go test -cover ./...
	go vet ./...

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

push:
	make test
	git pull origin HEAD
	git push origin HEAD

.DEFAULT_GOAL := test
