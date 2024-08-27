# First stage: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o cmd/app/app ./cmd/app

# Second stage: Create a minimal image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/cmd/app/app .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
