.PHONY: build-server build-fe build clean test-server run-server dev-fe

build-server:
	@echo "Building Backend"
	go build -o main ./cmd/main.go

build-fe:
	@echo "Building Frontend"
	@echo "Linting..."
	@bun run lint
	@echo "Formatting..."
	@cd frontend && bun run format
	@echo "Building..."
	@bun install --frozen-lockfile && bun run build

build: build-server build-fe

clean:
	@echo "Cleaning..."
	@rm -f main

test-server:
	@echo "Testing server..."
	@go test ./... -v

run-server:
	@go run ./backend/cmd/main.go

dev-fe:
	cd frontend && bun run dev
