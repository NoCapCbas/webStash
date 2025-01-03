dev-up:
	docker compose --env-file .env -f docker-compose.dev.yml up --build -d

dev-recreate:
	docker compose --env-file .env -f docker-compose.dev.yml up --build -d --force-recreate

prod:
	docker compose --env-file .env -f docker-compose.prod.yml up --build -d --force-recreate

dev-recreate:
	docker compose --env-file .env -f docker-compose.dev.yml up --build -d


help:
	@echo "Usage: make [command]"
	@echo "Commands:"
	@echo "  dev     Run development environment"
	@echo "  prod    Run production environment"

.PHONY: dev prod help

