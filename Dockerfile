# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Update module path
RUN go mod edit -module todo-app

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo-app

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/todo-app .
# Copy templates and other necessary files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/database ./database

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./todo-app"] 