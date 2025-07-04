dev:
	docker compose --env-file .env -f docker-compose.dev.yml up --build -d --force-recreate

prod:
	docker compose --env-file .env -f docker-compose.prod.yml up --build -d --force-recreate

test:
	docker build -f dockerfile.dev -t webstash-python-dev .
	docker run --rm -p 8080:8080 webstash-python-dev

test-prod:
	docker build -f dockerfile.prod -t webstash-python-prod .
	docker run --rm -p 8080:8080 webstash-python-prod
