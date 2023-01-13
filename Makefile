PROJECT_NAME := service
DOCKER_COMPOSE_PATH?=deployments/docker-compose.yml

#export PATH="$PATH:$(go env GOPATH)/bin"

start:
	docker compose -f $(DOCKER_COMPOSE_PATH) build --no-cache
	docker compose -f $(DOCKER_COMPOSE_PATH) up -d

start-integration-test:
	docker compose -f $(DOCKER_COMPOSE_PATH) build --no-cache
	docker compose -f $(DOCKER_COMPOSE_PATH) up -d
	docker compose -f $(DOCKER_COMPOSE_PATH) run integration_test /bin/sh -c "go test -cover ./..."; \
	test_status_code=$$?
	docker compose -f $(DOCKER_COMPOSE_PATH) down; \
	exit $$test_status_code

start-unit-test:
	go test -v -cover -gcflags=-l --race `go list ./... | grep -v /integration_test`

dep:
	@go mod download

install_linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint --version

run_linter:
	golangci-lint run -v --timeout 4m

mock-install:
	#GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0 (Если используешь Go version < 1.16)
	go install github.com/golang/mock/mockgen@v1.6.0

mock-gen:
	mockgen \
	-source ./internal/usecase/usecase.go \
	-destination ./internal/usecase/mock_usecase/usecase_mocks.go

build:
	CGO_ENABLED=0 go build -o ./bin/$(PROJECT_NAME) ./cmd/main.go