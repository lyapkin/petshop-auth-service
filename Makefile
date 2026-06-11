include .env
export $(shell sed 's/=.*//' .env)

.PHONY: certs

certs:
	mkdir -p $(SECRET_DIR)
	openssl genrsa -out $(JWT_ACCESS_SECRET_PATH) 2048
	openssl rsa -in $(JWT_ACCESS_SECRET_PATH) -pubout -out $(JWT_ACCESS_PUBLIC_PATH)
	@chmod 600 $(JWT_ACCESS_SECRET_PATH)
	@chmod 644 $(JWT_ACCESS_PUBLIC_PATH)
	@echo "Ключи успешно созданы в $(SECRET_DIR)"

clean-certs:
	rm -rf $(SECRET_DIR)

export GOOSE_DRIVER=${DB_DRIVER}
export GOOSE_DBSTRING=$(DB_URL)
export GOOSE_MIGRATION_DIR=$(MIGRATIONS_DIR)

DB_URL=${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
MIGRATIONS_DIR=./migrations

create:
	@goose create $(name) sql

up:
	@goose up

down:
	@goose down

status:
	@goose status