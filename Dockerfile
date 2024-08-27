# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o main core-customer/src/main.go

# Final Stage
FROM alpine:latest

WORKDIR /root/

# Copy the built application binary from the builder stage
COPY --from=builder /app/main .

# Copy the migrations directory from the builder stage
COPY --from=builder /app/core-customer/src/migrations /migrations

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
