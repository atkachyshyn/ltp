.PHONY: build run docker-build docker-run test clean

# Build the executable
build:
	@echo "Building the application..."
	@go build -o main ./cmd

# Run the application
run: build
	@echo "Running the application..."
	@./main

# Build the Docker container
docker-build:
	@echo "Building the Docker container..."
	@docker build -t bitcoin-ltp .

# Run the Docker container
docker-run: docker-build
	@echo "Running the Docker container..."
	@docker run -p 8080:8080 bitcoin-ltp

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean up the project
clean:
	@echo "Cleaning up..."
	@rm -f main
