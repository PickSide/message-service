run: build
	@./bin/message-service

build:
	@echo "Building Message service..."
	@go build -o ./bin/message-service ./cmd/main.go
	@echo "Done"