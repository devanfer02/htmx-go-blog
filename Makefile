# variables
DB_URL = postgres://postgres:postgres@localhost:5432/htmxblog?sslmode=disable

.PHONY: help
help:
	@echo "Choose a command:"
	@echo "  make run            - Run the Go application"
	@echo "  make migrate-up     - Apply all migrations"
	@echo "  make migrate-down   - Rollback all migrations"
	@echo "  make create-migration name=<name> - Create a new migration file"

.PHONY: run
run:
	go run ./cmd/go-blog/

.PHONY: migrate-up
migrate-up:
	migrate -path ./data/db/migrations -database ${DB_URL} up

.PHONY: migrate-down
migrate-down:
	migrate -path ./data/db/migrations -database ${DB_URL} down 

.PHONY: migrate-force
migrate-force:
ifndef version 
	$(error "Migration version not specified. Use 'make migrate-force version=<version>'")
endif 
	migrate -path ./data/db/migrations -database ${DB_URL} force $(version)

.PHONY: migrate-create
migrate-create:
ifndef name 
	$(error "Migration name not specified. Use 'make create-migration name=<name>'")
endif 
	migrate create -ext sql -dir ./migrations -seq $(name)