.PHONY: run-b
run-b: ## run local
	CONFIG_PATH=config/configb.yaml go run cmd/main.go


# ---------- help ----------
.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ {printf "  \033[36m%-24s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


# ---------- run ----------

.PHONY: compose-up
compose-up: ## up docker compose
	docker compose -f docker-compose.yaml up --build

.PHONY: compose-down
compose-down: ## down docker compose
	docker compose -f docker-compose.yaml down


# ---------- run local ----------

.PHONY: compose-up-local
compose-up-local: ## up docker compose local
	docker compose -f docker-compose.local.yaml up -d

.PHONY: compose-down-local
compose-down-local: ## down docker compose local
	docker compose -f docker-compose.local.yaml down

.PHONY: run
run: ## run local
	CONFIG_PATH=config/config.local.yaml go run cmd/main.go

.PHONY: run-test
run-test: ## run local for tests
	CONFIG_PATH=config/config_test.local.yaml go run cmd/main.go

..PHONY: build
build: ## build local
	go build -o main cmd/main.go


# ---------- migrations ----------

POSTGRES_SETUP := user=postgres password=postgres dbname=postgres host=localhost port=5433 sslmode=disable
POSTGRES_SETUP_TEST := user=postgres password=postgres dbname=test host=localhost port=5431 sslmode=disable
MIGRATION_FOLDER=$(CURDIR)/migrations
MIGRATION_NAME=pvz

.PHONY: migration-create
migration-create: ## create new migration
	goose -dir "$(MIGRATION_FOLDER)" create "$(MIGRATION_NAME)" sql

.PHONY: migration-up
migration-up: ## migration up
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down: ## migration down
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

.PHONY: migration-up-test
migration-up-test: ## migration up test
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: migration-down-test
migration-down-test: ## migration down test
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down


# ---------- tests ----------

.PHONY: test
test: ## run tests
	go test -v -count=2 -p=3 ./...

.PHONY: test-integration
test-integration: ## run integration tests
	CONFIG_PATH=$(CURDIR)/config/config_test.yaml go test ./... -tags=integration -v

.PHONY: test-colored
test-colored: ## run tests colored
	grc go test -v -count=2 -p=3 ./...

.PHONY: test-integration-colored
test-integration-colored: ## run integration tests colored
	CONFIG_PATH=$(CURDIR)/config/config_test.yaml grc go test ./... -tags=integration -v


# ---------- rare ----------

.PHONY: generate-proto
generate-proto: ## generate proto
	rm -rf internal/pkg swagger && \
	mkdir -p internal/pkg swagger
	protoc  --proto_path=internal/api_proto \
			--go_out=internal/pkg \
			--go-grpc_out=internal/pkg \
			--grpc-gateway_out=internal/pkg --grpc-gateway_opt generate_unbound_methods=true \
			--openapiv2_out=swagger \
 			api.proto

.PHONY: linter
linter: ## check by golangci linter
	golangci-lint run

.PHONY: docker-rm-volume
docker-rm-volume: ## remove docker volume
	docker volume rm pg-data

.PHONY: gen-ssl-cert
gen-ssl-cert: ## generate fresh ssl certificate
	openssl genrsa -out server.key 2048  # Сгенерировать приватный ключ (.key)
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365 -nodes  # Сгенерировать публичный ключ (.crt), но основе приватного
	mv -f server.key server.crt internal/server/certs/  # Поместить оба файла в папку /certificate
