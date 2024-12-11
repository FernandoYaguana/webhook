# Use the official Golang image to build the Go application
FROM golang:1.19 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage to create a smaller image for the final container
FROM debian:bullseye-slim

# Install necessary dependencies for running Go applications
RUN apt-get update && apt-get install -y ca-certificates

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main /usr/local/bin/main

# Expose the port the app will run on
EXPOSE 8080

# Run the Go application
CMD ["main"]