# Build the application
all: build test

build:
	@echo "Building ..."

	@source ~/.envs/.go_lin_env
	
	@go build -o ./bin/server ./cmd/server/main.go

win-build:
	@echo "Building ..."

	@source ~/.envs/.go_win_env

	@go build -buildmode=exe -o ./bin/win-server.exe ./cmd/server/main.go

# Run the application
run:
	@go run cmd/server/main.go

run-prod:
	./bin/main

# Create DB container
docker-run:
	@if docker compose up -d 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker compose up -d; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integrations Test for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	rm -f ./bin/server

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

sqlc:
	@sqlc generate

migrate-up:
	@GOOSE_DRIVER=sqlite3 GOOSE_MIGRATION_DIR=./sql/sqlite/schema GOOSE_DBSTRING=./server.db goose up  

migrate-down:
	@GOOSE_DRIVER=sqlite3 GOOSE_MIGRATION_DIR=./sql/sqlite/schema GOOSE_DBSTRING=./server.db goose down

.PHONY: all build win-build run test clean watch docker-run docker-down sqlc migrate-up migrate-down itest
