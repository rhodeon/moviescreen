SHELL := /bin/bash

include .env

# --- HELPERS ---
## help: display this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# confirm: display confirmation prompt
.PHONY: confirm
confirm:
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

# --- DEVELOPMENT ---
## run/api: run the cmd/api application
current_time = $(shell date +"%Y-%m-%dT%H:%M:%S%z")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = "-s -X main.buildTime=${current_time} -X main.version=${git_description}"

.PHONY: run/api
run/api:
	go run -ldflags=${linker_flags} ./cmd/api

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo "building cmd/api..."
	# build for local machine
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api

	# build for Linux
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api

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

# --- QUALITY CONTROL ---
## audit: tidy and vendor dependencies and format, vet and test codebase
.PHONY: audit
audit: vendor
	@echo "formatting codebase..."
	go fmt ./...

	@echo "vetting code..."
	go vet ./...
	staticcheck ./...

	@echo "running tests..."
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo "tidying and verifying module dependencies..."
	go mod tidy
	go mod verify

	@echo "vendoring dependencies..."
	go mod vendor

# --- PRODUCTION ---
remote = ${PRODUCTION_USER}@${PRODUCTION_HOST_IP}
remote_dir = ${remote}:~/service/

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	ssh -i ${PRIVATE_KEY_PATH} ${remote}

## production/deploy/api: deploy api build
.PHONY: production/deploy/api
production/deploy/api: build/api
	scp -i ${PRIVATE_KEY_PATH} ./bin/linux_amd64/api ${remote_dir}

# production/deploy/env: deploy production dotenv file
.PHONY: production/deploy/env
production/deploy/env:
	scp -i ${PRIVATE_KEY_PATH} ./remote/production/.env ${remote_dir}

## production/migrations: deploy and execute database migrations
.PHONY: production/migrations
production/migrations:
	scp -i ${PRIVATE_KEY_PATH} -r ./migrations ${remote_dir}
	ssh -t -i ${PRIVATE_KEY_PATH} ${remote} 'migrate -path ~/migrations -database $$MOVIESCREEN_DB_DSN up'

## production/configure/api.service: configure the production systemd api.service file
.PHONY: production/configure/api.service
production/configure/api.service:
	scp -i ${PRIVATE_KEY_PATH} -r ./remote/production/api.service ${remote_dir}
	ssh -t -i ${PRIVATE_KEY_PATH} ${remote} '\
	sudo mv	~/service/api.service /etc/systemd/system/ \
	&& sudo systemctl enable api \
	&& sudo systemctl restart api \
	'

## production/configure/caddyfile: configure the production Caddyfile
.PHONY: production/configure/caddyfile
production/configure/caddyfile:
	scp -i ${PRIVATE_KEY_PATH} -r ./remote/production/Caddyfile ${remote_dir}
	ssh -t -i ${PRIVATE_KEY_PATH} ${remote} '\
	sudo mv	~/service/Caddyfile /etc/caddy/ \
	&& sudo systemctl reload caddy \
	'