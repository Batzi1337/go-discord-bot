# Use the official Golang image to create a build artifact.
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o message-bot ./cmd/message-bot

# Start a new stage from scratch
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/message-bot .

# Copy the config file
# COPY internal/config/config.yaml ./internal/config/config.yaml

# Expose port (if your bot needs to listen on a port)
# EXPOSE 8080

# Command to run the executable
CMD ["./message-bot"]