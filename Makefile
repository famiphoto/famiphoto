# Makefile


start:
	docker compose -f docker-compose.yaml --env-file ./.env build --no-cache
	docker compose -f docker-compose.yaml --env-file ./.env up  -d

restart:
	docker compose -f docker-compose.yaml stop
	docker compose -f docker-compose.yaml up -d

dev:
	docker compose -f docker-compose.local.yaml --env-file ./.env build --no-cache
	docker compose -f docker-compose.local.yaml --env-file ./.env up  -d
