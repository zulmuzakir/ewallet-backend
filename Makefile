.PHONY: all build run test clean lint sqlc migrate-up migrate-down tidy

# Variables
APP_NAME := ewallet-backend
MAIN_PATH := ./cmd/api
BUILD_DIR := ./bin
MIGRATION_PATH := ./db/migrations
DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/db_ewallet?sslmode=disable

# Default target
all: tidy lint sqlc build

# Build the application
build:
	@echo "Building..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Run the application
run:
	@go run $(MAIN_PATH)

# Run tests
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...

# Run tests with coverage report
test-coverage: test
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

# Generate sqlc code
sqlc:
	@echo "Generating sqlc code..."
	@sqlc generate

# Run database migrations up
migrate-up:
	@echo "Running migrations..."
	@migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" up

# Run database migrations down
migrate-down:
	@echo "Rolling back migrations..."
	@migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" down

# Create a new migration
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $$name

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download

# Install development tools
tools:
	@echo "Installing development tools..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/api/main.go -o docs

# Format Swagger annotations
swagger-fmt:
	@echo "Formatting Swagger annotations..."
	@swag fmt

# Help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  lint          - Run linter"
	@echo "  sqlc          - Generate sqlc code"
	@echo "  migrate-up    - Run database migrations"
	@echo "  migrate-down  - Rollback database migrations"
	@echo "  migrate-create- Create a new migration"
	@echo "  tidy          - Tidy dependencies"
	@echo "  deps          - Download dependencies"
	@echo "  tools         - Install development tools"
