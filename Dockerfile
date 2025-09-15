# Use the official Golang image as the build environment
FROM golang:1.23-alpine AS builder

# Install git (needed for fetching modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o url-shortner ./main.go

# -----------------------------
# Final lightweight image
# -----------------------------
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/url-shortner .

COPY .env .

# Expose the service port
EXPOSE 8080

# Run the application
CMD ["./url-shortner"]