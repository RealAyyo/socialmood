include .env
BIN := "./bin/socialmood"
DOCKER_IMG="socialmood:develop"

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/crash

run: build
	$(BIN) -config ./configs/config.yaml

migrate:
	goose -dir ./migrations postgres "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up


rollback:
	goose -dir ./migrations postgres "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down