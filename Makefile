
build_be:
	mkdir -p bin
	cd backend && go build -o ../bin/flashcards ./cmd/
run_be:
	cd backend && go run ./cmd/
dev_fe:
	cd frontend && bun run dev
