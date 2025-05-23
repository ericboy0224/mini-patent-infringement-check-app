.PHONY: build run stop clean test logs help frontend-install

# Default target
.DEFAULT_GOAL := help

# Variables
DOCKER_COMPOSE = docker compose
APP_NAME = patlytics

# Help command
help:
	@echo "Available commands:"
	@echo "  make build            - Build Docker images"
	@echo "  make run              - Run the application in detached mode"
	@echo "  make stop             - Stop all containers"
	@echo "  make clean            - Stop and remove containers, networks, and volumes"
	@echo "  make test             - Run tests"
	@echo "  make logs             - View application logs"
	@echo "  make restart          - Restart the application"
	@echo "  make frontend-install - Install frontend dependencies"

# Build Docker images
build:
	$(DOCKER_COMPOSE) build

# Run the application
run:
	$(DOCKER_COMPOSE) up -d

# Stop containers
stop:
	$(DOCKER_COMPOSE) stop

# Clean up everything
clean:
	$(DOCKER_COMPOSE) down -v
	docker system prune -f
	cd frontend && rm -rf node_modules && rm -rf dist

# Run tests
test:
	go test -v ./...
	cd frontend && pnpm test

# View logs
logs:
	$(DOCKER_COMPOSE) logs -f

# Restart application
restart: stop run

# Frontend commands
frontend-install:
	cd frontend && pnpm install

# One-command setup for first time run
setup:
	@echo "Installing frontend dependencies..."
	$(MAKE) frontend-install
	@echo "Building Docker images..."
	$(MAKE) build
	@echo "Starting services..."
	$(MAKE) run
	@echo "Waiting for services to start..."
	@sleep 5
	@echo "Application is running at http://localhost:8080"

# Development commands
dev-build:
	$(MAKE) frontend-install
	$(MAKE) build
	$(MAKE) run
	$(MAKE) logs
