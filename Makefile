all: lint

.PHONY: lint
lint:
	golangci-lint run --fix --fast
