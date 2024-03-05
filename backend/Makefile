BINPATH ?= build

.PHONY: lint
lint:
	golangci-lint run ./... --timeout 2m --tests=false --skip-dirs=features

.PHONY: test
test:
	go test -race -cover ./...


