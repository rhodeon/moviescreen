SHELL := /bin/bash

include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# confirm: display confirmation prompt
.PHONY: confirm
confirm:
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api

## db/migrations/create name=$1: create a new database migration
.PHONY: db/migrations/create
db/migrations/create: confirm
	@echo "creating migration files for ${name}..."
	migrate create -seq -ext .sql -dir ./migrations ${name}

## db/migrations/up: display current database migration version
.PHONY: db/migrations/version
db/migrations/version:
	@echo -n "database migration version: "
	@migrate -path ./migrations -database ${DB_DSN} version

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo "running up migrations..."
	@migrate -path ./migrations -database ${DB_DSN} up

## db/migrations/down: rollback all database migrations
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo "running down migrations..."
	@migrate -path ./migrations -database ${DB_DSN} down

## db/migrations/goto version=$1: move to a specified database migration version
.PHONY: db/migrations/goto
db/migrations/goto: confirm
	@echo "migrating database to version ${version}..."
	@migrate -path ./migrations -database ${DB_DSN} goto ${version}