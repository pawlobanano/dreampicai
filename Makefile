DB_URL=postgresql://$(PROJECT_NAME)_local_container_postgresql_database_username:$(PROJECT_NAME)_local_container_postgresql_database_password@localhost:5432/$(PROJECT_NAME)_local_container_postgresql_database?sslmode=disable
PROJECT_NAME=dreampicai

build:
	templ generate view
	go build -tags develop -o bin/$(PROJECT_NAME) main.go

create-local-container-network:
	podman network create $(PROJECT_NAME)_local_container_network

create-local-container-postgresql:
	podman run --name $(PROJECT_NAME)_local_container_postgresql --network $(PROJECT_NAME)_local_container_network --mount type=volume,source=$(PROJECT_NAME)_local_container_postgresql_database_volume,target=/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_USER=$(PROJECT_NAME)_local_container_postgresql_database_username -e POSTGRES_PASSWORD=$(PROJECT_NAME)_local_container_postgresql_database_password -d postgres:16.3-alpine3.19

create-local-container-postgresql-database:
	podman exec -it $(PROJECT_NAME)_local_container_postgresql createdb --username=$(PROJECT_NAME)_local_container_postgresql_database_username --owner=$(PROJECT_NAME)_local_container_postgresql_database_username $(PROJECT_NAME)_local_container_postgresql_database

drop-local-postgresql-database:
	podman exec -it $(PROJECT_NAME)_local_container_postgresql dropdb --username=$(PROJECT_NAME)_local_container_postgresql_database_username $(PROJECT_NAME)_local_container_postgresql_database

install:
	go install github.com/a-h/templ/cmd/templ@latest
	go get ./...
	go mod vendor
	go mod tidy
	go mod download
	npm install -D daisyui@latest

lint:
	golangci-lint run ./...

migrate-down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrate-down-1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

migrate-up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrate-up-1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

run: build
	./bin/$(PROJECT_NAME)

.PHONY: build create-local-container-network create-local-container-postgresql create-local-container-postgresql-database drop-local-postgresql-database install lint migrate-down migrate-down-1 migrate-up migrate-up-1 run
