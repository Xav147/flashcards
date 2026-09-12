.PHONY: help build-server build-fe build clean test-server run-server dev-fe

help: # Show list of make targets
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

build-server: # Build the Go backend and produce the `main` binary
	@echo "Building Backend"
	go build -o main ./cmd/main.go

build-fe: # Lint, format, and build the frontend
	@echo "Building Frontend"
	@echo "Linting..."
	@bun run lint
	@echo "Formatting..."
	@cd frontend && bun run format
	@echo "Building..."
	@bun install --frozen-lockfile && bun run build

build: build-server build-fe # Build both the server and frontend

clean: # Remove the compiled `main` binary
	@echo "Cleaning..."
	@rm -f main

test-server: # Run the Go server tests verbosely
	@echo "Testing server..."
	@go test ./... -v

run-server: # Run the backend server
	@go run ./backend/cmd/main.go

dev-fe: # Start the frontend dev server
	cd frontend && bun run dev
