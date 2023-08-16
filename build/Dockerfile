# Use a multi-stage build for a smaller final image
# This stage builds the Go application
FROM golang:1.19 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the server binary
WORKDIR /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o server

# This stage uses a smaller base image for the final runtime container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binaries from the builder stage
COPY --from=builder /app/cmd/server/server .

# Expose the ports the server will use
EXPOSE 50051

# Command to run the server binary
CMD ["./server"]