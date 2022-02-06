PG_CONTAINER ?= learngo-pg
PG_USER ?= postgres
PG_PASSWORD ?= postgres
PG_DB ?= bookshop
PG_PORT ?= 5433
PG_URL ?= "postgresql://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_DB)?sslmode=disable"
PG_MIGRATIONS_DIR = "./migrations/postgres"

.PHONY: pgrun
pgrun:
	docker run --name $(PG_CONTAINER) -e POSTGRES_USER=$(PG_USER) -e POSTGRES_PASSWORD=$(PG_PASSWORD) \
	-e POSTGRES_DB=$(PG_DB) -p $(PG_PORT):5432 -d postgres

.PHONY: pgc
pgc:
	docker exec -it $(PG_CONTAINER) bash -c "PGPASSWORD=$(PG_PASSWORD) psql -U $(PG_USER) $(PG_DB)"

.PHONY: deps
deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: run
run:
	PG_URL=$(PG_URL) go run ./cmd/main.go

test:
	PG_TEST_DB_URL=$(PG_URL) go test -v -race ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: pgstatus
pgstatus:
	 goose -dir $(PG_MIGRATIONS_DIR) postgres $(PG_URL) status

.PHONY: pgup
pgup:
	 goose -dir $(PG_MIGRATIONS_DIR) postgres $(PG_URL) up

.PHONY: pgdown
pgdown:
	 goose -dir $(PG_MIGRATIONS_DIR) postgres $(PG_URL) down
