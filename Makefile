.PHONY: start-dev-rebuild start-dev stop-dev log-services

start-dev-rebuild:
	docker compose -f docker-compose.dev.yml up -d --build

start-dev:
	docker compose -f docker-compose.dev.yml up -d

stop-dev:
	docker compose -f docker-compose.dev.yml down

log-services:
	docker compose -f docker-compose.dev.yml logs -f $(s)

.PHONY: migrate-create migrate-up migrate-down migrate-status

migrate-create:
	goose create $(name) sql

migrate-up:
	goose up

migrate-down:
	goose down

migrate-status:
	goose status

.PHONY: gen-swagger

gen-swagger:
	swag init -d ./cmd/api/ 