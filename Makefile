.PHONY: start-dev-rebuild start-dev stop-dev log-services

start-dev-rebuild:
	docker compose -f docker-compose.dev.yml up -d --build

start-dev:
	docker compose -f docker-compose.dev.yml up -d

stop-dev:
	docker compose -f docker-compose.dev.yml down

log-services:
	docker compose -f docker-compose.dev.yml logs -f $(s)