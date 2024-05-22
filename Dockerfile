FROM golang:1.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application code to the working directory
COPY src/ ./src/

# Set the working directory to src
WORKDIR /app/src

# Build the Go application
RUN go build -o /app/main .

# Stage 2: Create a minimal image for the final executable
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
