setup.golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINARY=golangci-lint bash -s -- v1.55.2

.PHONY: lint
lint: setup.golangci-lint
	./bin/golangci-lint run

.PHONY: build
build:
	go build -o ./bin/ ./cmd/...

