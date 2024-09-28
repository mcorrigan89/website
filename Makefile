ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif

.PHONY: start
start:
	./bin/main

.PHONY: build
build:
	go build -o=./bin/main ./cmd

.PHONY: codegen
codegen:
	git submodule update --recursive --remote  
	buf lint
	buf generate --path serviceapis/serviceapis

.PHONY: models
models:
	pg_dump --schema-only website_builder > schema.sql
	sqlc generate

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	migrate -path=./migrations -database="$(POSTGRES_URL)?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	migrate -path=./migrations -database="$(POSTGRES_URL)?sslmode=disable" down 1
