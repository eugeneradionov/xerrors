export GO111MODULE=on

.PHONY: dep lint test

dep: ## Download required dependencies
	go mod vendor
	go mod tidy

lint: ## Lint files
	golangci-lint run -c .golangci.yml

test: dep ## Run unit tests
	go test -cover -race -count=1 ./...
