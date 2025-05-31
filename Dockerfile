# Build stage
FROM golang:1.22-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory that matches Go module path
WORKDIR /go/src/github.com/Shokh-web/go-todo-list

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/todo-app

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /go/bin/todo-app .
# Copy templates and other necessary files
COPY --from=builder /go/src/github.com/Shokh-web/go-todo-list/templates ./templates
COPY --from=builder /go/src/github.com/Shokh-web/go-todo-list/database ./database

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./todo-app"]