swagger:
	swag init --parseDependency --parseInternal -g /server/server.go

compose:
	docker compose -f docker-compose.yaml up -d
