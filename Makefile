APP_NAME = "prometheus-metric-validator"

PHONY: build
build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) ./cmd/main.go

PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...
