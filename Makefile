dev:
	docker compose -f docker-compose.dev.yml up --build -d

prod:
	docker compose -f docker-compose.prod.yml up --build -d
