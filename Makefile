dev:
	docker compose --env-file .env -f docker-compose.dev.yml up --build

dev-recreate:
	docker compose --env-file .env -f docker-compose.dev.yml up --build -d --force-recreate

dev-logs:
	docker compose --env-file .env -f docker-compose.dev.yml logs

dev-exec:
	docker exec -it webstash-users-1 bash

dev-down:
	docker compose -f docker-compose.dev.yml down

prod:
	docker compose --env-file .env -f docker-compose.prod.yml up --build -d --force-recreate
