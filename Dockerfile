# Use a newer Go version that satisfies the requirement
FROM golang:1.23.4-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Install dependencies for Alpine Linux
RUN apk add --no-cache git postgresql-client

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire application code
COPY . .

# Copy the .env file into the container
COPY .env .env

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Set the entry point for the container
CMD ["./main"]
