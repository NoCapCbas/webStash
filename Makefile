dev-recreate:
	docker compose --env-file .env -f docker-compose.dev.yml up --build -d --force-recreate

dev-up:
	docker compose --env-file .env -f docker-compose.dev.yml up --build -d

prod:
	docker compose --env-file .env -f docker-compose.prod.yml up --build -d --force-recreate
