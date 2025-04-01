include .env

migration_path=internal/storage/postgres/migrations

create_migration:
	migrate create -ext=sql -dir=${migration_path} $(migration_name)

migrate_up:
	migrate -path=$(migration_path) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=$(migration_path) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" -verbose down

.PHONY: create_migration migrate_up migrate_down