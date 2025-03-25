# Stage 1: Build the application
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go module files separately to leverage caching
COPY go.mod go.sum ./

# Download dependencies (cache layer)
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /melvin .

# Stage 2: Create a lightweight production image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /melvin /app/melvin

# Ensure the binary is executable
RUN chmod +x /app/melvin

# Set the entrypoint and default command
ENTRYPOINT ["/app/melvin"]
CMD ["-config", "/var/lib/config.toml"]
