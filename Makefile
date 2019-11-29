BINARY_NAME=algorithm_web_service

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: mod-download
mod-download: ## download modules to local cache
	go mod download

.PHONY: build
build: mod-download ## build binary for local OS
	go build -o $(BINARY_NAME)

.PHONY: build-linux
build-linux: mod-download ## build linux binary
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)_linux

.PHONY: run
run: build ## build and run app locally
	./$(BINARY_NAME)

.PHONY: test
test: ## run tests
	go test -v

.PHONY: bench
bench: ## run benchmark and unittests
	go test -bench=.

.PHONY: clean
clean: ## run benchmark and unittests
	rm -rf $(BINARY_NAME)*

.PHONY: compose-up
compose-up: compose-down compose-build ## Create and start containers
	docker-compose up -d

.PHONY: compose-down
compose-down: ## Stop and remove containers, networks, images, and volumes
	docker-compose down --remove-orphans

.PHONY: compose-restart
compose-restart: compose-up ## restart services

.PHONY: compose-tail
compose-tail: ## Tail output from containers
	docker-compose logs -f

.PHONY: compose-build
compose-build: ## Build or rebuild services
	docker-compose build --no-cache

.PHONY: compose-top
compose-top: ## Display the running processes
	docker-compose top

.PHONY: compose-ps
compose-ps: ## List containers
	docker-compose ps

