# Makefile

## Start Famiphoto components for hosting
start:
	docker compose -f docker-compose.yaml --env-file ./.env build --no-cache
	docker compose -f docker-compose.yaml --env-file ./.env up  -d

## Restart Famiphoto components
restart:
	docker compose -f docker-compose.yaml stop
	docker compose -f docker-compose.yaml up -d

## Start Famiphoto components as development mode
dev:
	docker compose -f docker-compose.local.yaml --env-file ./.env build --no-cache
	docker compose -f docker-compose.local.yaml --env-file ./.env up  -d

## Enter API container
dc_exec_api:
	docker compose -f docker-compose.local.yaml exec famiphoto_api bash
