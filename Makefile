# Makefile for building and running the Go project

.DEFAULT_GOAL := all

.PHONY: install build run clean test deps docker-build docker-run docker-clean all

# Install dependencies
install: clean
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# Build the project
build: clean
	go build -o cmd/app/app ./cmd/app

# Run the project
run:
	cmd/app/app

# Clean the project
clean:
	go clean
	rm -f swagger.json
	rm -f cmd/app/app
	rm -f *.tar.gz

lint:
	golangci-lint --verbose run ./...

# Test the project
test:
	go test -v ./...

# Install dependencies
deps:
	go get -u ./...

swagger:
	swagger generate spec -o ./swagger.json --scan-models --spec=3.0 --input cmd/app/main.go

# Download and unpack swagger-ui
download-swagger-ui:
	@if [ -f "swagger-ui.tar.gz" ]; then \
		rm swagger-ui.tar.gz; \
	fi
	@if [ ! -d "./swagger-ui" ]; then \
		curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest \
		| grep "tarball_url" \
		| cut -d '"' -f 4 \
		| xargs curl -L -o swagger-ui.tar.gz; \
		mkdir -p swagger-ui-temp; \
		tar -xzf swagger-ui.tar.gz --strip-components=1 -C ./swagger-ui-temp; \
		mkdir -p swagger-ui; \
		cp -r ./swagger-ui-temp/dist/* ./swagger-ui/; \
		rm -rf ./swagger-ui-temp; \
		rm swagger-ui.tar.gz; \
	fi

# Build Docker image
docker-build: swagger
	docker build -t go-reference-api:latest .

# Run Docker container
docker-run:
	docker run -p 8080:8080 go-reference-api:latest

# Remove Docker image
docker-clean:
	docker rmi go-reference-api:latest

# Default target
all: build
