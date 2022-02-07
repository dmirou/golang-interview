PG_CONTAINER ?= learngo-pg
PG_USER ?= postgres
PG_PASSWORD ?= postgres
PG_DB ?= bookshop
PG_PORT ?= 5433
PG_URL ?= "postgresql://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_DB)?sslmode=disable"
PG_MIGRATIONS_DIR = "./migrations/postgres"

PG_ENV = -e "POSTGRES_USER=$(PG_USER)" -e "POSTGRES_PASSWORD=$(PG_PASSWORD)" \
	-e "POSTGRES_DB=$(PG_DB)"

PG_DB_PATH = ~/docker/bookshop
PG_CONTAINER_DB_PATH = /var/lib/postgresql/data

PG_RUNNED := $(shell docker ps -f name=^/$(PG_CONTAINER)$$ | grep $(PG_CONTAINER))
PG_EXITED := $(shell docker ps -a -f name=^/$(PG_CONTAINER)$$ | grep $(PG_CONTAINER))

.PHONY: pgstart
pgstart:
ifndef PG_RUNNED
ifndef PG_EXITED
		mkdir -p $(PG_DB_PATH)
		docker run -d --name $(PG_CONTAINER) \
		-p $(PG_PORT):5432 -v $(PG_DB_PATH):$(PG_CONTAINER_DB_PATH) \
		$(PG_ENV) postgres
else
		docker start $(PG_CONTAINER)
endif
endif

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
	PG_TEST_DB_URL=$(PG_URL) go test -count=1 -v -race ./...

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
