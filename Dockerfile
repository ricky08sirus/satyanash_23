# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o webserver .

# Final stage
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/webserver /app/webserver

# Set the working directory
WORKDIR /app

COPY .env /app/.env

# Expose the application port (if your app runs on port 8080)
EXPOSE 8080

# Run the application
CMD ["./webserver"]

