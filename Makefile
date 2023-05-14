lint:
	@echo "--> Running linter"
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
	@golangci-lint run --timeout=10m

build:
	@echo "--> Building app"
	go build -o app cmd/main.go

.PHONY: lint build