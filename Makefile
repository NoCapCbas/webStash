watch-dev:
	docker compose --env-file .env.dev watch

watch-prod:
	docker compose --env-file .env.prod watch

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  watch-dev: Watch the development environment"
	@echo "  watch-prod: Watch the production environment"
	@echo "  help: Display this help message"

.PHONY: watch-dev watch-prod help
