.DEFAULT_GOAL := help

.PHONY: install
install: ## Initial setup
	mise install
	go mod tidy

.PHONY: lint
lint:  ## Lint go files
	golangci-lint run --config=./.golangci.yml ./...

.PHONY: fmt
fmt:  ## Format go files
	go fmt ./...

.PHONY: test
test: ## Run tests
	go test ./...

.PHONY: run
run: ## Run server
	go run ./cmd/server/main.go

.PHONY: generate-mocks
generate-mocks: ## Generate mocks files
	go generate ./...

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
